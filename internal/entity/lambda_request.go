package entity

type LambdaRequest struct {
	RequestID    string  `json:"request_id"`
	SQSMessageID string  `json:"message_id"`
	ErrCode      *string `json:"code"`
	ErrMessage   *string `json:"message"`
}
