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
	QueueMap               map[string]*domain.Queue
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

	// TODO_Done_Modified : using locks to ensure concurrent operations can be carried out
	// locks on QueueManager rather than QueueMap due to deadlock problem
	qm.TransactionLockManager.AcquireLock([]string{"QueueManager"})
	qm.QueueMap[request.Name] = &domain.Queue{
		ID: request.Name,
	}
	qm.TransactionLockManager.ReleaseLock([]string{"QueueManager"})
	logger.Logger.Info().Msgf("Created Queue : %s", request.Name)

	return nil
}
func (qm *QueueManager) GetQueue(queueName string) (*domain.Queue, error) {
	// maps are not good for concurrent reads
	// TODO_Done_Modified: we want to acquire a lock on the queue-map
	// locks on QueueManager rather than QueueMap due to deadlock problem
	qm.TransactionLockManager.AcquireLock([]string{"QueueManager"})
	queueRef := qm.QueueMap[queueName]
	qm.TransactionLockManager.ReleaseLock([]string{"QueueManager"})
	if queueRef == nil {
		return nil, QueueDoesNotExists
	}

	return queueRef, nil
}
