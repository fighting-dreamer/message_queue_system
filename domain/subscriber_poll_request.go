package domain

type SubscriberPollRequest struct {
	SubscriberID string `json:"subscriber_id"`
	MessageID    int    `json:"message_id"`
	FetchCount   int    `json:"fetch_count"`
}
