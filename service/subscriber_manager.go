package service

type ISubscriberManager interface {
	RegisterSubscriber()
	ValidateSubscriberRequest()
	GetQueueSubscribers()
}
