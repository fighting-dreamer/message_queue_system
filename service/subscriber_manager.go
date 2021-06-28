package service

import "nipun.io/message_queue/domain"

type ISubscriberManager interface {
	RegisterSubscriber(request *domain.SubscriberRegisterRequest) error
	GetQueueSubscribers(queueName string) []*domain.Subscriber
	GetSubscriberQueueName(subscriberID string) string
	IncrementUnackCounter(queueName string, subscriberID string) error
	IncrementAckCounter(queueName string, subscriberID string) error
	GetUnackCounter(queueName string, subscriberID string) int64
	GetAckCounter(queueName string, subscriberID string) int64
}
