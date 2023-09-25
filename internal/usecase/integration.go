package usecase

import (
	"aws-lambda-sqs/internal/entity"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type integrationUseCase struct {
	masterRepo MasterRepository
}

// NewIntegrationUseCase
func NewIntegrationUseCase(mRepo MasterRepository,
) (IntegrationUseCase, error) {
	return &integrationUseCase{
		masterRepo: mRepo,
	}, nil
}

// Integration
func (iuc *integrationUseCase) Integration(sqsEvent events.SQSEvent) error {
	var lambdaReq entity.LambdaRequest
	lambdaRequests := make([]entity.LambdaRequest, len(sqsEvent.Records))
	for i, message := range sqsEvent.Records {
		// Decode JSON
		err := json.Unmarshal([]byte(message.Body), &lambdaReq)
		if err != nil {
			return err
		}

		// SQSEventをLambdaRequestに格納
		lambdaRequests[i] = entity.LambdaRequest{
			RequestID:    lambdaReq.RequestID,
			SQSMessageID: message.MessageId,
			ErrCode:      lambdaReq.ErrCode,
			ErrMessage:   lambdaReq.ErrMessage,
		}
	}
	// Log書き込み
	err := iuc.masterRepo.WriteLog(lambdaRequests)
	if err != nil {
		return err
	}
	return nil
}
