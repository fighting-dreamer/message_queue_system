package domain

type SubscriberRegisterResponse struct {
	SubscriberID string `json:"subscriber_id"`
	QueueName    string `json:"queue_name"`
}
