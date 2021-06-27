package domain

type SubscriberPollResponse struct {
	Messages []Message `json:"messages"`
}
