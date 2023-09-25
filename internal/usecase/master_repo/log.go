package master_repo

import (
	"aws-lambda-sqs/internal/entity"
	"fmt"

	"github.com/google/uuid"
)

type Logs struct {
	LogID        string  `gorm:"column:log_id"`
	RequestID    string  `gorm:"column:request_id"`
	SQSMessageID string  `gorm:"column:sqs_message_id"`
	ErrCode      *string `gorm:"column:code"`
	ErrMessage   *string `gorm:"column:message"`
}

func (Logs) TableName() string {
	return "logs"
}

func (r *MasterRepository) WriteLog(
	lambdaRequests []entity.LambdaRequest,
) error {
	madeLogs := makeCreateLog(lambdaRequests)
	fmt.Println("Log LogID scheduled to be stored :", madeLogs[0].LogID)
	fmt.Println("Log RequestID scheduled to be stored :", madeLogs[0].RequestID)
	fmt.Println("Log SQSMessageID scheduled to be stored :", madeLogs[0].SQSMessageID)
	fmt.Println("Log ErrCode scheduled to be stored :", &madeLogs[0].ErrCode)
	fmt.Println("Log ErrMessage scheduled to be stored :", &madeLogs[0].ErrMessage)
	return nil
}

func makeCreateLog(
	lambdaRequests []entity.LambdaRequest,
) []Logs {
	logs := make([]Logs, len(lambdaRequests))
	for i, lr := range lambdaRequests {
		logs[i] = Logs{
			LogID:        uuid.New().String(),
			RequestID:    lr.RequestID,
			SQSMessageID: lr.SQSMessageID,
			ErrCode:      lr.ErrCode,
			ErrMessage:   lr.ErrMessage,
		}
	}
	return logs
}
