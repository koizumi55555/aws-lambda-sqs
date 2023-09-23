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
	if err := r.DBHandler.Conn.Create(&madeLogs).Error; err != nil {
		return fmt.Errorf(
			"db connection error or unknown db error: %s", err)
	}
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
