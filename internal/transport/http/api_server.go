package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kapustaprusta/promotions-service/v2/internal/config"
	"github.com/kapustaprusta/promotions-service/v2/internal/domain"
	"github.com/kapustaprusta/promotions-service/v2/internal/repository"
	"github.com/kapustaprusta/promotions-service/v2/internal/transport"
)

// PromotionService declares interface for promotion services
type PromotionService interface {
	TruncateAll(context.Context) error
	Insert(context.Context, *domain.PromotionModel) error
	FindByRecordID(context.Context, int) (*domain.PromotionModel, error)
}

type apiServer struct {
	router           *mux.Router
	promotionService PromotionService
}

// Start launches api server
func Start(config *config.Config, promotionService PromotionService) error {
	return http.ListenAndServe(config.BindAddr, newAPIServer(promotionService))
}

func newAPIServer(promotionService PromotionService) *apiServer {
	s := &apiServer{
		router:           mux.NewRouter(),
		promotionService: promotionService,
	}

	s.configureRouter()

	return s
}

// ServeHTTP serves http requests
func (s *apiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *apiServer) configureRouter() {
	s.router.HandleFunc("/promotions/{record_id}", s.handlePromotionsFind()).Methods("GET")
	s.router.HandleFunc("/promotions", s.handlePromotionsUpload()).Methods("POST")
}

func (s *apiServer) handlePromotionsFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		recordID, err := strconv.Atoi(vars["record_id"])

		if err != nil {
			s.badRequest("cannot-parse-record-id", err, w, r)
			return
		}

		promotion, err := s.promotionService.FindByRecordID(r.Context(), recordID)
		if err != nil {
			if errors.Is(err, repository.ErrRecordNotFound) {
				s.notFound("promotion-not-found", err, w, r)
				return
			}

			s.respondWithError(err, w, r)
			return
		}

		s.respondOK(transport.DomainModel2TransportModel(promotion), w, r)
	}
}

func (s *apiServer) handlePromotionsUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.promotionService.TruncateAll(r.Context())
		if err != nil {
			s.respondWithError(err, w, r)
			return
		}

		log.Println("starts uploading promotions")

		errCh := make(chan error)
		doneCh := make(chan struct{})
		promotionCh := make(chan *transport.PromotionModel)

		defer func() {
			close(errCh)
			close(doneCh)
			close(promotionCh)
		}()

		go func() {
			err := transport.ReadPromotionsStream(r.Context(), r.Body, promotionCh)
			if err != nil {
				errCh <- err
			} else {
				doneCh <- struct{}{}
			}
		}()

		promotionsCount := 0
		for {
			select {
			case <-r.Context().Done():
				log.Printf("request cancelled")
				return
			case <-doneCh:
				log.Printf("finises uploading promotions")
				s.respondOK(map[string]int{"total_promotions": promotionsCount}, w, r)
				return
			case err := <-errCh:
				log.Printf("error while parsing promotion csv: %+v", err)
				s.badRequest("invalid-csv", err, w, r)
				return
			case transportModel := <-promotionCh:
				promotionsCount++
				log.Printf("[%d] received promotion: %+v", promotionsCount, *transportModel)

				domainModel, err := transport.TransportModel2DomainModel(transportModel)
				if err != nil {
					s.badRequest("invalid-input", err, w, r)
					return
				}

				if err := s.promotionService.Insert(r.Context(), domainModel); err != nil {
					s.respondWithError(err, w, r)
					return
				}
			}
		}
	}
}
