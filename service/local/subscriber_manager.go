package service

import (
	"errors"
	"nipun.io/message_queue/domain"
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
}

// TODO : Create Subscriber

func (sm *SubscriberManager) RegisterSubscriber(request *domain.SubscriberRegisterRequest) error {
	// check if queue exists
	_, err := sm.QueueManager.GetQueue(request.QueueName)
	if err != nil {
		return err
	}
	// check if subscriber exist and is already registered to that queue
	queueName := sm.SubscriberToQueueMap[request.SubscriberID]
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
	// TODO : using locks to ensure concurrent operations can be carried out
	sm.SubscriberMap[request.SubscriberID] = &domain.Subscriber{
		ID:  request.SubscriberID,
		URL: request.URL,
	}
	sm.SubscriberToQueueMap[request.SubscriberID] = request.QueueName
	sm.QueueToSubscriberListMap[request.QueueName] = append(sm.QueueToSubscriberListMap[request.QueueName], sm.SubscriberMap[request.SubscriberID])
	return nil
}

func (sm *SubscriberManager) GetQueueSubscribers(queueName string) []*domain.Subscriber {
	return sm.QueueToSubscriberListMap[queueName]
}

func (sm *SubscriberManager) GetSubscriberQueueName(subscriberID string) string {
	return sm.SubscriberToQueueMap[subscriberID]
}
