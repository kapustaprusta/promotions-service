package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kapustaprusta/promotions-service/v2/internal/domain"
	"github.com/kapustaprusta/promotions-service/v2/internal/services"
)

var _ services.PromotionRepository = &InMemPromotionRepository{}

// InMemPromotionRepository declares in-memory promotion repository
type InMemPromotionRepository struct {
	rwMutex    *sync.RWMutex
	promotions map[int]*domain.PromotionModel
}

// NewInMemPromotionRepository creates new in-memory promotion repository
func NewInMemPromotionRepository() *InMemPromotionRepository {
	return &InMemPromotionRepository{
		rwMutex:    &sync.RWMutex{},
		promotions: make(map[int]*domain.PromotionModel),
	}
}

// TruncateAll removes all recorde from repository that are older than ttl
func (s *InMemPromotionRepository) TruncateAll(_ context.Context) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	s.promotions = make(map[int]*domain.PromotionModel)

	return nil
}

// Insert inserts promotion into repository
func (s *InMemPromotionRepository) Insert(_ context.Context, promotion *domain.PromotionModel) error {
	if promotion == nil {
		return fmt.Errorf(
			"%w: promotion cannot be nil",
			ErrNilValue,
		)
	}

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	nextRecordID := len(s.promotions)
	_ = promotion.SetCreatedAt(time.Now())
	s.promotions[nextRecordID] = promotion

	return nil
}

// FindByRecordID finds promotion in repository by record ID
func (s *InMemPromotionRepository) FindByRecordID(_ context.Context, recordID int) (*domain.PromotionModel, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	promotionModel, isFound := s.promotions[recordID]
	if !isFound {
		return nil, fmt.Errorf(
			"%w: cannot find promotion by ID '%d'",
			ErrRecordNotFound, recordID,
		)
	}

	return promotionModel, nil
}
