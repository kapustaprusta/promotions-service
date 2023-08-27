package domain

import "time"

// PromotionModel declares promotion domain model
type PromotionModel struct {
	id             string
	price          float64
	createdAt      time.Time
	expirationDate time.Time
}

// NewPromotionModel creates new promotion domain model
func NewPromotionModel(
	id string,
	price float64,
	createdAt time.Time,
	expirationDate time.Time,
) (*PromotionModel, error) {
	if err := validateInput(id, price, createdAt, expirationDate); err != nil {
		return nil, err
	}

	return &PromotionModel{
		id:             id,
		price:          price,
		createdAt:      createdAt,
		expirationDate: expirationDate,
	}, nil
}

// ID returns promotion id
func (m *PromotionModel) ID() string {
	return m.id
}

// SetID sets promotion id
func (m *PromotionModel) SetID(id string) error {
	if err := validateID(id); err != nil {
		return nil
	}

	m.id = id

	return nil
}

// Price returns promotion price
func (m *PromotionModel) Price() float64 {
	return m.price
}

// SetPrice sets promotion price
func (m *PromotionModel) SetPrice(price float64) error {
	if err := validatePrice(price); err != nil {
		return nil
	}

	m.price = price

	return nil
}

// CreatedAt returns promotion created at
func (m *PromotionModel) CreatedAt() time.Time {
	return m.createdAt
}

// SetCreatedAt sets promotion created at
func (m *PromotionModel) SetCreatedAt(createdAt time.Time) error {
	if err := validateCreatedAt(createdAt); err != nil {
		return err
	}

	m.createdAt = createdAt

	return nil
}

// ExpirationDate returns promotion expiration date
func (m *PromotionModel) ExpirationDate() time.Time {
	return m.expirationDate
}

// SetExpirationDate sets promotion expiration date
func (m *PromotionModel) SetExpirationDate(expirationDate time.Time) error {
	if err := validateExpirationDate(expirationDate); err != nil {
		return nil
	}

	m.expirationDate = expirationDate

	return nil
}
