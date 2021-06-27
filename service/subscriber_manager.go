package service

import "nipun.io/message_queue/domain"

type ISubscriberManager interface {
	RegisterSubscriber(request *domain.SubscriberRegisterRequest) error
	GetQueueSubscribers(queueName string) []*domain.Subscriber
}
