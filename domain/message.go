package domain

type Metadata struct {
	QueueName string `json:"queue_name"`
}
type Message struct {
	ID       string      `json:"id"`
	Metadata Metadata    `json:"metadata"`
	Value    interface{} `json:"value"`
}
