package main

import (
	db "aws-lambda-sqs/db/postgres"
	"aws-lambda-sqs/internal/controller"
	"aws-lambda-sqs/internal/usecase"
	"aws-lambda-sqs/internal/usecase/master_repo"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	masterDBH, err := db.NewDBHandler()
	if err != nil {
		log.Fatalf("DBHandler error: %s", err)
	}

	// repo
	mRepo := master_repo.New(masterDBH)
	// usecase
	integrationUC, err := usecase.NewIntegrationUseCase(mRepo)
	if err != nil {
		log.Fatal("init usecase error", err)
		return
	}
	h := controller.NewSqsHandler(integrationUC)

	lambda.Start(h.SqsHandler)
}
