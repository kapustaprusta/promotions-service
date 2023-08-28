package transport

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05 -0700 MST"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeLayout))
	return []byte(stamp), nil
}

// ReadPromotionsCsvStream reads promotion from provided reader and sends them to chan.
func ReadPromotionsCsvStream(ctx context.Context, reader io.Reader, outputCh chan<- *PromotionModel) error {
	csvReader := csv.NewReader(reader)
	for {
		// Check if context is cancelled.
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// Read line from csv
		promotionProps, err := csvReader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return fmt.Errorf("failed to read promotion from source: %w", err)
		}

		if len(promotionProps) != 3 {
			return errors.New("failed to parse promotion: too few properties")
		}

		// Read promotion and send it to the channel.
		promotionModel := &PromotionModel{}
		promotionModel.ID = promotionProps[0]

		promotionModel.Price, err = strconv.ParseFloat(promotionProps[1], 64)
		if err != nil {
			return fmt.Errorf("failed to parse promotion: invalid price: %w", err)
		}

		promotionExpirationDate, err := time.Parse(timeLayout, promotionProps[2])
		if err != nil {
			return fmt.Errorf("failed to parse promotion: invalid expiration date: %w", err)
		}
		promotionModel.ExpirationDate = JSONTime(promotionExpirationDate)

		outputCh <- promotionModel
	}

	return nil
}
