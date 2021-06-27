package domain

type SubscriberRegisterRequest struct {
	// Metadata or other auth info not taken right now.
	SubscriberID string `json:"subscriber_id"`
	QueueName    string `json:"queue_name"`
}
