package transport

import (
	"time"

	"github.com/kapustaprusta/promotions-service/v2/internal/domain"
)

// PromotionModel declares promotion transport model
type PromotionModel struct {
	ID             string
	Price          float64
	CreatedAt      time.Time
	ExpirationDate time.Time
}

// DomainModel2TransportModel converts domain model to transport model
func DomainModel2TransportModel(domainModel *domain.PromotionModel) *PromotionModel {
	return &PromotionModel{
		ID:             domainModel.ID(),
		Price:          domainModel.Price(),
		ExpirationDate: domainModel.ExpirationDate(),
	}
}

// TransportModel2DomainModel converts transport model to domain model
func TransportModel2DomainModel(transportModel *PromotionModel) (*domain.PromotionModel, error) {
	return domain.NewPromotionModel(
		transportModel.ID,
		transportModel.Price,
		transportModel.CreatedAt,
		transportModel.ExpirationDate,
	)
}
