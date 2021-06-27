package service

import (
	"errors"
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
)

var (
	QueueDoesNotExists = errors.New("QueueDoesNotExists")
)

type QueueManager struct {
	QueueMap map[string]*domain.Queue
}

func (qm *QueueManager) CreateQueue(request domain.CreateQueueRequest) error {
	queueref, err := qm.GetQueue(request.Name)
	if err == nil && queueref != nil {
		// Queue already Exists
		// Idempotent Operation
		return nil
	}

	if err != QueueDoesNotExists {
		return err
	}

	// TODO : using locks to ensure concurrent operations can be carried out
	qm.QueueMap[request.Name] = &domain.Queue{
		ID: request.Name,
	}
	logger.Logger.Info().Msgf("Created Queue : %s", request.Name)

	return nil
}
func (qm *QueueManager) GetQueue(queueName string) (*domain.Queue, error) {
	queueRef := qm.QueueMap[queueName]
	if queueRef == nil {
		return nil, QueueDoesNotExists
	}

	return queueRef, nil
}
