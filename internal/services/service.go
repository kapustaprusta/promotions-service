package services

import (
	"context"
	"github.com/kapustaprusta/promotions-service/v2/internal/domain"
)

// PromotionRepository declares interface for promotion repository
type PromotionRepository interface {
	TruncateAll(context.Context) error
	Insert(context.Context, *domain.PromotionModel) error
	FindByRecordID(context.Context, int) (*domain.PromotionModel, error)
}

// PromotionService declares promotion services
type PromotionService struct {
	repository PromotionRepository
}

// NewPromotionService create new promotion services
func NewPromotionService(repository PromotionRepository) *PromotionService {
	return &PromotionService{
		repository: repository,
	}
}

// TruncateAll removes all records from repository
func (s *PromotionService) TruncateAll(ctx context.Context) error {
	return s.repository.TruncateAll(ctx)
}

// Insert inserts promotion into repository
func (s *PromotionService) Insert(ctx context.Context, promotion *domain.PromotionModel) error {
	return s.repository.Insert(ctx, promotion)
}

// FindByRecordID finds promotion by record ID in repository
func (s *PromotionService) FindByRecordID(ctx context.Context, id int) (*domain.PromotionModel, error) {
	return s.repository.FindByRecordID(ctx, id)
}
