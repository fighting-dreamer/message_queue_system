package service

import (
	"errors"
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

var (
	RegisteredToOtherQueue = errors.New("RegisteredToOtherQueue")
)

type SubscriberManager struct {
	// It is simplified for as subscriber and queue ID are just strings.
	SubscriberMap            map[string]*domain.Subscriber
	SubscriberToQueueMap     map[string]string
	QueueToSubscriberListMap map[string][]*domain.Subscriber
	QueueManager             service.IQueueManager
	TransactionLockManager   service.ITransactionLockManager
}

// TODO : Create Subscriber

func (sm *SubscriberManager) RegisterSubscriber(request *domain.SubscriberRegisterRequest) error {
	// check if queue exists
	_, err := sm.QueueManager.GetQueue(request.QueueName)
	if err != nil {
		return err
	}
	// check if subscriber exist and is already registered to that queue
	// TODO_Done_Modified : use lock for the use of SubscriberToQueueMap
	// locks on SubscriberManager rather than SubscriberToQueueMap, due to deadlock problem
	sm.TransactionLockManager.AcquireLock([]string{"SubscriberManager"})
	queueName := sm.SubscriberToQueueMap[request.SubscriberID]
	sm.TransactionLockManager.ReleaseLock([]string{"SubscriberManager"})

	// TODO : validation check for queue registration to not have empty string queues
	if queueName != "" {
		// subscriber is registered to a queue already.
		if queueName == request.QueueName {
			// idempotent API
			return nil
		} else {
			return RegisteredToOtherQueue
		}
	}
	// subscriber is new and is not registered to any queue
	// TODO_Done_Modified : using locks on SubscriberMap to ensure concurrent operations can be carried out
	// TODO_Done_Modified : using locks on SubscriberToQueueMap to ensure concurrent operations can be carried out
	// TODO_Done_Modified : using locks on QueueToSubscriberListMap to ensure concurrent operations can be carried out
	// locks on SubscriberManager rather than SubscriberToQueueMap,SubscriberMap,QueueToSubscriberListMap ; due to deadlock problem
	sm.TransactionLockManager.AcquireLock([]string{"SubscriberManager"})

	sm.SubscriberMap[request.SubscriberID] = &domain.Subscriber{
		ID:  request.SubscriberID,
		URL: request.URL,
	}
	sm.SubscriberToQueueMap[request.SubscriberID] = request.QueueName
	sm.QueueToSubscriberListMap[request.QueueName] = append(sm.QueueToSubscriberListMap[request.QueueName], sm.SubscriberMap[request.SubscriberID])

	sm.TransactionLockManager.ReleaseLock([]string{"SubscriberManager"})
	return nil
}

func (sm *SubscriberManager) GetQueueSubscribers(queueName string) []*domain.Subscriber {
	// TODO_Done_Modified : using locks on QueueToSubscriberListMap to ensure concurrent operations can be carried out
	// locks on SubscriberManager rather than QueueToSubscriberListMap, due to deadlock problem
	sm.TransactionLockManager.AcquireLock([]string{"SubscriberManager"})
	subsribers := sm.QueueToSubscriberListMap[queueName]
	logger.Logger.Debug().Msgf("subscribers for queue %s : %+v", queueName, subsribers)
	sm.TransactionLockManager.ReleaseLock([]string{"SubscriberManager"})
	return subsribers
}

func (sm *SubscriberManager) GetSubscriberQueueName(subscriberID string) string {
	// TODO_Done_Modified : using locks on SubscriberToQueueMap to ensure concurrent operations can be carried out
	// locks on SubscriberManager rather than SubscriberToQueueMap, due to deadlock problem
	sm.TransactionLockManager.AcquireLock([]string{"SubscriberManager"})
	queueName := sm.SubscriberToQueueMap[subscriberID]
	sm.TransactionLockManager.ReleaseLock([]string{"SubscriberManager"})
	return queueName
}
