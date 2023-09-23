package controller

import (
	"aws-lambda-sqs/internal/usecase"
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type SqsHandler struct {
	IntegrationUsecase usecase.IntegrationUseCase
}

// NewSqsHandler
func NewSqsHandler(uc usecase.IntegrationUseCase) SqsHandler {
	return SqsHandler{
		IntegrationUsecase: uc,
	}
}

// SqsHandler
func (h SqsHandler) SqsHandler(ctx context.Context, sqsEvent events.SQSEvent) error {
	// integration呼び出し
	if len(sqsEvent.Records) < 1 {
		return nil
	}
	return h.IntegrationUsecase.Integration(sqsEvent)
}
