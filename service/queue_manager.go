package service

import "nipun.io/message_queue/domain"

type IQueueManager interface {
	CreateQueue()
	GetQueue(queueName string) (*domain.Queue, error)
}
