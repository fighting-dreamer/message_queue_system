package domain

type AckMessageRequest struct {
	MessageID    int64  `json:"message_id"`
	QueueName    string `json:"queue_name"`
	SubscriberID string `json:"subscriber_id"`
}
