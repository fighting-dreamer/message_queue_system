package domain

type SubscriberRegisterRequest struct {
	// Metadata or other auth info not taken right now.
	SubscriberID string `json:"subscriber_id"`
	URL          string `json:"callback_url"`
	QueueName    string `json:"queue_name"`
}
