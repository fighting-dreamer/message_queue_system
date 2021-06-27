package service

import "nipun.io/message_queue/domain"

type SubscriberManager struct {
}

func (sm *SubscriberManager) RegisterSubscriber(request *domain.SubscriberRegisterRequest) error {
	return nil
}
func (sm *SubscriberManager) ValidateSubscriberRequest(request *domain.SubscriberRegisterRequest) error {
	return nil
}
func (sm *SubscriberManager) GetQueueSubscribers() {}
