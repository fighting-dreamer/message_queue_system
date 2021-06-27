package domain

type SubscriberRegisterRequest struct {
	// Metadata or other auth info not taken right now.
	ID        string
	QueueName string
}
