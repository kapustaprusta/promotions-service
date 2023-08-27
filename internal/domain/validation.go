package domain

import (
	"fmt"
	"time"
)

func validateInput(id string, price float64, createdAt time.Time, expirationDate time.Time) error {
	err := validatePromotionID(id)
	if err != nil {
		return err
	}

	err = validatePromotionPrice(price)
	if err != nil {
		return err
	}

	err = validatePromotionCreatedAt(createdAt)
	if err != nil {
		return err
	}

	// TODO: generate valid test data
	//err = validatePromotionExpirationDate(expirationDate)
	//if err != nil {
	//	return err
	//}

	return nil
}

func validatePromotionID(id string) error {
	if len(id) == 0 {
		return fmt.Errorf(
			"%w: promotion ID cannot be empty",
			ErrInvalidValue,
		)
	}

	return nil
}

func validatePromotionPrice(price float64) error {
	if price < 0.0 {
		return fmt.Errorf(
			"%w: promotion price cannot be less than zero",
			ErrInvalidValue,
		)
	}

	return nil
}

func validatePromotionCreatedAt(createdAt time.Time) error {
	if createdAt.UnixMilli() > time.Now().UnixMilli() {
		return fmt.Errorf(
			"%w: created at  cannot be greate than current time",
			ErrInvalidValue,
		)
	}

	return nil
}

func validatePromotionExpirationDate(expirationDate time.Time) error {
	if expirationDate.UnixMilli() <= time.Now().UnixMilli() {
		return fmt.Errorf(
			"%w: promotion expiration date cannot be less than current time",
			ErrInvalidValue,
		)
	}

	return nil
}
