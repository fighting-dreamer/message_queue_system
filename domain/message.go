package domain

type Metadata struct {
	QueueName string
}
type Message struct {
	ID       string
	Metadata Metadata
	Value    interface{}
}
