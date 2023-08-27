package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	testPromotionID             = "d018ef0b-dbd9-48f1-ac1a-eb4d90e57118"
	testPromotionPrice          = 60.683466
	testPromotionCreatedAt      = time.Now().Add(time.Minute * (-10))
	testPromotionExpirationDate = time.Now().Add(time.Minute * 10)
)

func TestPromotionModel_NewPromotion(t *testing.T) {
	t.Parallel()

	t.Run(
		"valid",
		func(t *testing.T) {
			promotion, err := NewPromotionModel(
				testPromotionID,
				testPromotionPrice,
				testPromotionCreatedAt,
				testPromotionExpirationDate,
			)

			require.NoError(t, err)
			require.Equal(t, testPromotionID, promotion.ID())
			require.Equal(t, testPromotionPrice, promotion.Price())
			require.Equal(t, testPromotionCreatedAt, promotion.CreatedAt())
			require.Equal(t, testPromotionExpirationDate, promotion.ExpirationDate())
		},
	)

	t.Run(
		"invalid promotion ID",
		func(t *testing.T) {
			_, err := NewPromotionModel(
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
			_, err := NewPromotionModel(
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
			_, err := NewPromotionModel(
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
	//		_, err := NewPromotionModel(
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
