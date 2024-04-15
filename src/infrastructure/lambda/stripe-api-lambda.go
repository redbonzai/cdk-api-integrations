package lambda

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go/charge"
)

type PaymentRequest struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	SourceToken string `json:"source"`
	Description string `json:"description"`
}

type PaymentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func StripeHandler(ctx context.Context, request PaymentRequest) (*PaymentResponse, error) {
	response, err := createCharge(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func createCharge(request PaymentRequest) (*PaymentResponse, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(request.Amount),
		Currency:    stripe.String(request.Currency),
		Description: stripe.String(request.Description),
	}
	params.SetSource(request.SourceToken)

	ch, err := charge.New(params)
	if err != nil {
		return &PaymentResponse{Success: false, Message: err.Error()}, nil
	}

	return &PaymentResponse{Success: true, Message: fmt.Sprintf("Charge created successfully: %s", ch.ID)}, nil
}

func main() {
	lambda.Start(StripeHandler)
}
