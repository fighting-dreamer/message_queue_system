package service

import "nipun.io/message_queue/domain"

type MessageBrokerService struct {
}

func (mbs *MessageBrokerService) SetMessage(queue *domain.Queue, message *domain.Message) (*domain.Message, error) {
	return nil, nil
}
func (mbs *MessageBrokerService) GetMessage()      {}
func (mbs *MessageBrokerService) DeleteMessage()   {}
func (mbs *MessageBrokerService) CallSubscribers() {}
