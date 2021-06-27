package domain

type Metadata struct {
}
type Message struct {
	ID       string
	Metadata Metadata
	Value    interface{}
}
