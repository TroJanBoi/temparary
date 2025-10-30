package usecases

import (
	"context"

	"github.com/TroJanBoi/temparary/internal/services/repository"
	"github.com/TroJanBoi/temparary/internal/services/types"
)

type PaymentUseCases interface {
	PaymentUseCase(ctx context.Context, pay *types.PaymentRequest) (*types.PaymentResponse, error)
}

type paymentUseCases struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentUseCases(paymentRepo repository.PaymentRepository) PaymentUseCases {
	return &paymentUseCases{
		paymentRepo: paymentRepo,
	}
}

func (p *paymentUseCases) PaymentUseCase(ctx context.Context, pay *types.PaymentRequest) (*types.PaymentResponse, error) {
	return p.paymentRepo.Payment(ctx, pay)
}
