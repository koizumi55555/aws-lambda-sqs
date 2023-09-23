package main

import (
	"aws-lambda-sqs/internal/controller"
	"aws-lambda-sqs/internal/usecase"
	"aws-lambda-sqs/internal/usecase/master_repo"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// repo
	mRepo := master_repo.New()
	// usecase
	integrationUC, err := usecase.NewIntegrationUseCase(mRepo)
	if err != nil {
		log.Fatal("init usecase error", err)
		return
	}
	h := controller.NewSqsHandler(integrationUC)

	lambda.Start(h.SqsHandler)
}
