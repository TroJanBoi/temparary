package controller

import (
	"log"
	"net/http"

	"github.com/TroJanBoi/temparary/internal/services/types"
	"github.com/TroJanBoi/temparary/internal/services/usecases"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	PaymentUseCase usecases.PaymentUseCases
}

func NewPayment(paymentUseCase usecases.PaymentUseCases) *PaymentController {
	return &PaymentController{
		PaymentUseCase: paymentUseCase,
	}
}

type TonowPayCallback struct {
	PaymentChannel         string `json:"paymentChannel"`
	DocumentId             string `json:"documentId"`
	DocumentNo             string `json:"documentNo"`
	CustomerName           string `json:"customerName"`
	CustomerEmail          string `json:"customerEmail"`
	CustomerPhone          string `json:"customerPhone"`
	ConfirmId              string `json:"confirmId"`
	PaymentId              string `json:"paymentId"`
	TransactionId          string `json:"transactionId"`
	TransactionDateandTime string `json:"transactionDateandTime"`
	TransactionAmount      string `json:"transactionAmount"`
	TransactionFee         string `json:"transactionFee"`
	TransactionTax         string `json:"transactionTax"`
	TransactionStatus      string `json:"transactionStatus"`
	TransactionName        string `json:"transactionName"`
	ReferenceId1           string `json:"referenceId1"`
	ReferenceId2           string `json:"referenceId2"`
	ReferenceId3           string `json:"referenceId3"`
	ReferenceId4           string `json:"referenceId4"`
	ReferenceId5           string `json:"referenceId5"`
}

// @Summary      Process Payment
// @Description  Process a payment request
// @Tags         payment
// @Accept       json
// @Produce      json
// @Param body body types.PaymentRequest true "Payment info"
// @Success      200   {object}  types.PaymentResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /payment [post]
func (p *PaymentController) PaymentController(ctx *gin.Context) {
	var request types.PaymentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	response, err := p.PaymentUseCase.PaymentUseCase(ctx, &request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to process payment"})
		return
	}

	ctx.JSON(200, response)
}

// @Summary Payment Callback (TonowPay)
// @Description Receive callback from TonowPay after payment success/failure
// @Tags payment
// @Accept json
// @Produce json
// @Param body body TonowPayCallback true "TonowPay callback payload"
// @Success 200 {object} map[string]string
// @Router /payment/callback [post]
func (p *PaymentController) PaymentCallback(ctx *gin.Context) {
	var payload TonowPayCallback
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	log.Printf("[TonowPayCallback] PaymentID=%s Status=%s Amount=%s",
		payload.PaymentId, payload.TransactionStatus, payload.TransactionAmount)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  payload.TransactionStatus,
		"message": "Callback accepted",
	})
}

func (p *PaymentController) PaymentRoutes(r gin.IRoutes) {
	r.POST("/", p.PaymentController)
	r.POST("/callback", p.PaymentCallback)
}
