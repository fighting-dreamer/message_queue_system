package domain

type CreateQueueRequest struct {
	// Metadata eg : is queue durable, just in-memory, replication etc.
	Name string `json:"name"`
}
