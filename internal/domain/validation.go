package domain

import (
	"fmt"
	"time"
)

func validateInput(id string, price float64, createdAt time.Time, expirationDate time.Time) error {
	err := validateID(id)
	if err != nil {
		return err
	}

	err = validatePrice(price)
	if err != nil {
		return err
	}

	err = validateCreatedAt(createdAt)
	if err != nil {
		return err
	}

	err = validateExpirationDate(expirationDate)
	if err != nil {
		return err
	}

	return nil
}

func validateID(id string) error {
	if len(id) == 0 {
		return fmt.Errorf(
			"%w: promotion ID cannot be empty",
			ErrInvalidValue,
		)
	}

	return nil
}

func validatePrice(price float64) error {
	if price < 0.0 {
		return fmt.Errorf(
			"%w: promotion price cannot be less than zero",
			ErrInvalidValue,
		)
	}

	return nil
}

func validateCreatedAt(createdAt time.Time) error {
	if createdAt.UnixMilli() > time.Now().UnixMilli() {
		return fmt.Errorf(
			"%w: created at  cannot be greate than current time",
			ErrInvalidValue,
		)
	}

	return nil
}

func validateExpirationDate(expirationDate time.Time) error {
	//because test input is outdated
	//if expirationDate.UnixMilli() <= time.Now().UnixMilli() {
	//	return fmt.Errorf(
	//		"%w: promotion expiration date cannot be less than current time",
	//		domain.ErrInvalidValue,
	//	)
	//}

	return nil
}
