package http_transport

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/kapustaprusta/promotions-service/v2/internal/transport"
)

type ErrorResponse struct {
	Slug       string `json:"slug"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}

func (s *apiServer) respondOK(response any, w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func (s *apiServer) respondWithError(err error, w http.ResponseWriter, r *http.Request) {
	var slugError transport.SlugError
	ok := errors.As(err, &slugError)
	if !ok {
		s.internalError("internal-server-error", err, w, r)
		return
	}

	switch slugError.ErrorType() {
	case transport.ErrorTypeIncorrectInput:
		s.badRequest(slugError.Slug(), slugError, w, r)
	case transport.ErrorTypeNotFound:
		s.notFound(slugError.Slug(), slugError, w, r)
	default:
		s.internalError(slugError.Slug(), slugError, w, r)
	}
}

func (s *apiServer) badRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	s.httpRespondWithError(err, slug, "Bad request", http.StatusBadRequest, w, r)
}

func (s *apiServer) notFound(slug string, err error, w http.ResponseWriter, r *http.Request) {
	s.httpRespondWithError(err, slug, "Not found", http.StatusNotFound, w, r)
}

func (s *apiServer) internalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	s.httpRespondWithError(err, slug, "Internal server error", http.StatusInternalServerError, w, r)
}

func (s *apiServer) httpRespondWithError(
	err error,
	slug string,
	msg string,
	status int,
	w http.ResponseWriter,
	_ *http.Request,
) {
	log.Printf("error: %s, slug: %s, msg: %s", err, slug, msg)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ErrorResponse{slug, status})
}
