package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/TroJanBoi/temparary/internal/services/types"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Payment(ctx context.Context, pay *types.PaymentRequest) (*types.PaymentResponse, error)
}

type paymentRepository struct {
	// add any necessary fields here
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (p *paymentRepository) Payment(ctx context.Context, pay *types.PaymentRequest) (*types.PaymentResponse, error) {
	endPoint := "https://gateway.tonow.net/v3/paymentToken"

	jsonData, err := json.Marshal(pay)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payment request: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", endPoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("transaction timeout: request exceeded 5 minutes")
		}
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("payment gateway returned status %d: %s", resp.StatusCode, string(body))
	}

	var paymentResponse types.PaymentResponse
	if err := json.Unmarshal(body, &paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payment response: %w", err)
	}

	return &paymentResponse, nil
}
