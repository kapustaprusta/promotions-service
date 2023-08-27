package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidate_ValidateInput(t *testing.T) {
	t.Parallel()

	t.Run(
		"valid",
		func(t *testing.T) {
			err := validateInput(
				testPromotionID,
				testPromotionPrice,
				testPromotionCreatedAt,
				testPromotionExpirationDate,
			)

			require.NoError(t, err)
		},
	)

	t.Run(
		"invalid promotion ID",
		func(t *testing.T) {
			err := validateInput(
				"",
				testPromotionPrice,
				testPromotionCreatedAt,
				testPromotionExpirationDate,
			)

			require.ErrorIs(t, err, ErrInvalidValue)
		},
	)

	t.Run(
		"invalid promotion price",
		func(t *testing.T) {
			err := validateInput(
				testPromotionID,
				-25.0,
				testPromotionCreatedAt,
				testPromotionExpirationDate,
			)

			require.ErrorIs(t, err, ErrInvalidValue)
		},
	)

	t.Run(
		"invalid promotion created at",
		func(t *testing.T) {
			err := validateInput(
				testPromotionID,
				testPromotionPrice,
				time.Now().Add(time.Hour),
				testPromotionExpirationDate,
			)

			require.ErrorIs(t, err, ErrInvalidValue)
		},
	)

	// TODO: generate valid test data
	//t.Run(
	//	"invalid promotion expiration date",
	//	func(t *testing.T) {
	//		err := validateInput(
	//			testPromotionID,
	//			testPromotionPrice,
	//			testPromotionCreatedAt,
	//			time.Now().Add(time.Hour*(-10)),
	//		)
	//
	//		require.ErrorIs(t, err, ErrInvalidValue)
	//	},
	//)
}

func TestValidate_ValidatePromotionID(t *testing.T) {
	t.Parallel()

	t.Run(
		"valid",
		func(t *testing.T) {
			require.NoError(t, validatePromotionID(testPromotionID))
		},
	)

	t.Run(
		"invalid promotion ID",
		func(t *testing.T) {
			require.ErrorIs(t, validatePromotionID(""), ErrInvalidValue)
		},
	)
}

func TestValidate_ValidatePromotionPrice(t *testing.T) {
	t.Parallel()

	t.Run(
		"valid",
		func(t *testing.T) {
			require.NoError(t, validatePromotionPrice(testPromotionPrice))
		},
	)

	t.Run(
		"invalid promotion price",
		func(t *testing.T) {
			require.ErrorIs(t, validatePromotionPrice(-25.0), ErrInvalidValue)
		},
	)
}

func TestValidate_ValidatePromotionCreatedAt(t *testing.T) {
	t.Parallel()

	t.Run(
		"valid",
		func(t *testing.T) {
			require.NoError(t, validatePromotionCreatedAt(testPromotionCreatedAt))
		},
	)

	t.Run(
		"invalid promotion created at",
		func(t *testing.T) {
			require.ErrorIs(t, validatePromotionCreatedAt(time.Now().Add(time.Hour)), ErrInvalidValue)
		},
	)
}

func TestValidate_ValidatePromotionExpirationDate(t *testing.T) {
	t.Parallel()

	t.Run(
		"valid",
		func(t *testing.T) {
			require.NoError(t, validatePromotionExpirationDate(testPromotionExpirationDate))
		},
	)

	t.Run(
		"invalid promotion expiration date",
		func(t *testing.T) {
			require.ErrorIs(t, validatePromotionExpirationDate(time.Now().Add(time.Hour*(-10))), ErrInvalidValue)
		},
	)
}
