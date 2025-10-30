package types

type PaymentRequest struct {
	AccountID      string                 `json:"accountId"`
	AccountChannel string                 `json:"accountChannel"`
	Amount         string                 `json:"amount"`
	Detail         string                 `json:"detail"`
	CustomerEmail  string                 `json:"customerEmail"`
	CustomerName   string                 `json:"customerName"`
	IsSMS          bool                   `json:"isSMS"`
	ReferenceID1   string                 `json:"referenceId1"`
	ReferenceID2   string                 `json:"referenceId2"`
	ReferenceID3   string                 `json:"referenceId3"`
	ReferenceID4   string                 `json:"referenceId4"`
	BackgroundURL  string                 `json:"backgroundUrl"`
	RedirectURL    string                 `json:"redirectUrl"`
	Charge         map[string]interface{} `json:"charge"`
}

type PaymentResponse struct {
	ID        string                 `json:"id"`
	Code      string                 `json:"code"`
	Status    string                 `json:"status"`
	Message   string                 `json:"message"`
	Timestamp int64                  `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}
