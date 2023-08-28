package transport

import (
	"time"

	"github.com/kapustaprusta/promotions-service/v2/internal/domain"
)

// PromotionModel declares promotion transport model
type PromotionModel struct {
	ID             string    `json:"id"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"-"`
	ExpirationDate JSONTime  `json:"expiration_date"`
}

// DomainModel2TransportModel converts domain model to transport model
func DomainModel2TransportModel(domainModel *domain.PromotionModel) *PromotionModel {
	return &PromotionModel{
		ID:             domainModel.ID(),
		Price:          domainModel.Price(),
		ExpirationDate: JSONTime(domainModel.ExpirationDate()),
	}
}

// TransportModel2DomainModel converts transport model to domain model
func TransportModel2DomainModel(transportModel *PromotionModel) (*domain.PromotionModel, error) {
	return domain.NewPromotionModel(
		transportModel.ID,
		transportModel.Price,
		transportModel.CreatedAt,
		time.Time(transportModel.ExpirationDate),
	)
}
