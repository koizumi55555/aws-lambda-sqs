package usecase

import (
	"aws-lambda-sqs/internal/entity"

	"github.com/aws/aws-lambda-go/events"
)

type (
	MasterRepository interface {
		// log.go
		WriteLog(lambdaReq []entity.LambdaRequest) error
	}

	IntegrationUseCase interface {
		Integration(sqsEvent events.SQSEvent) error
	}
)
