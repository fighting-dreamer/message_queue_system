package service

import "nipun.io/message_queue/domain"

type IQueueManager interface {
	CreateQueue(request domain.CreateQueueRequest) error
	GetQueue(queueName string) (*domain.Queue, error)
}
