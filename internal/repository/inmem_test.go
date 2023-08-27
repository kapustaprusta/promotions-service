package repository

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kapustaprusta/promotions-service/v2/internal/domain"
	"github.com/stretchr/testify/require"
)

var (
	testPromotionPrice          = 60.683466
	testPromotionCreatedAt      = time.Now().Add(time.Minute * (-10))
	testPromotionExpirationDate = time.Now().Add(time.Minute * 10)
)

func TestPromotionRepository_CreateOrUpdatePort(t *testing.T) {
	t.Parallel()
	repository := NewInMemPromotionRepository()

	t.Run(
		"insert promotion",
		func(t *testing.T) {
			t.Parallel()

			expectedPromotion := newRandomPromotion(t)

			recordID, err := repository.Insert(context.Background(), expectedPromotion)
			require.NoError(t, err)

			actualPromotion, err := repository.FindByRecordID(context.Background(), recordID)
			require.NoError(t, err)
			require.Equal(t, expectedPromotion, actualPromotion)
		},
	)

	t.Run(
		"find by id promotion",
		func(t *testing.T) {
			t.Parallel()

			expectedPromotion := newRandomPromotion(t)

			recordID, err := repository.Insert(context.Background(), expectedPromotion)
			require.NoError(t, err)

			actualPromotion, err := repository.FindByRecordID(context.Background(), recordID)
			require.NoError(t, err)
			require.Equal(t, expectedPromotion, actualPromotion)

			actualPromotion, err = repository.FindByRecordID(context.Background(), recordID+10)
			require.ErrorIs(t, err, ErrRecordNotFound)
			require.Nil(t, actualPromotion)
		},
	)

	t.Run(
		"truncate all promotions",
		func(t *testing.T) {
			t.Parallel()

			recordIDs := make([]int, 3)
			actualPromotions := make([]*domain.PromotionModel, 3)
			expectedPromotions := make([]*domain.PromotionModel, 3)

			for idx, _ := range expectedPromotions {
				expectedPromotions[idx] = newRandomPromotion(t)
				recordIDs[idx], _ = repository.Insert(
					context.Background(),
					expectedPromotions[idx],
				)
			}

			for idx, recordID := range recordIDs {
				actualPromotions[idx], _ = repository.FindByRecordID(
					context.Background(),
					recordID,
				)
			}

			require.Equal(t, expectedPromotions, actualPromotions)

			err := repository.TruncateAll(context.Background())
			require.NoError(t, err)

			for _, recordID := range recordIDs {
				_, err = repository.FindByRecordID(
					context.Background(),
					recordID,
				)

				require.ErrorIs(t, err, ErrRecordNotFound)
			}
		},
	)

	t.Run(
		"nil promotion",
		func(t *testing.T) {
			t.Parallel()

			_, err := repository.Insert(context.Background(), nil)
			require.ErrorIs(t, err, ErrNilValue)
		},
	)
}

func newRandomPromotion(t *testing.T) *domain.PromotionModel {
	t.Helper()

	randomPromotionID := uuid.New().String()
	promotion, err := domain.NewPromotionModel(
		randomPromotionID,
		testPromotionPrice,
		testPromotionCreatedAt,
		testPromotionExpirationDate,
	)
	require.NoError(t, err)

	return promotion
}
