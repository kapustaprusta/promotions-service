package transport

import (
	"context"

	"github.com/kapustaprusta/promotions-service/v2/internal/domain"
)

// PromotionService declares interface for promotion services
type PromotionService interface {
	TruncateAll(context.Context) error
	Insert(context.Context, *domain.PromotionModel) (int, error)
	FindByRecordID(context.Context, int) (*domain.PromotionModel, error)
}
