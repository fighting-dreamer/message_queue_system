package service

import (
	"errors"
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

var (
	QueueDoesNotExists = errors.New("QueueDoesNotExists")
)

type QueueManager struct {
	QueueMap map[string]*domain.Queue
	TransactionLockManager service.ITransactionLockManager
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
	qm.TransactionLockManager.AcquireLock([]string{"QueueMap"})
	qm.QueueMap[request.Name] = &domain.Queue{
		ID: request.Name,
	}
	qm.TransactionLockManager.ReleaseLock([]string{"QueueMap"})
	logger.Logger.Info().Msgf("Created Queue : %s", request.Name)

	return nil
}
func (qm *QueueManager) GetQueue(queueName string) (*domain.Queue, error) {
	// maps are not good for concurrent reads
	// TODO: we want to acquire a lock on the queue-map
	qm.TransactionLockManager.AcquireLock([]string{"QueueMap"})
	queueRef := qm.QueueMap[queueName]
	qm.TransactionLockManager.ReleaseLock([]string{"QueueMap"})
	if queueRef == nil {
		return nil, QueueDoesNotExists
	}

	return queueRef, nil
}
