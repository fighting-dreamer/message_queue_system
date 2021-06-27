package service

import "nipun.io/message_queue/domain"

type QueueManager struct {
}

func (qm *QueueManager) CreateQueue() {}
func (qm *QueueManager) GetQueue(queueName string) (*domain.Queue, error)    {
	return nil, nil
}
