package service

import (
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

type MessageBrokerService struct {
	MessageStoreService service.IMessageStoreService
	CallBackChan        chan *domain.Message
}

func (mbs *MessageBrokerService) SetMessage(queue *domain.Queue, message *domain.Message) (*domain.Message, error) {
	mbs.MessageStoreService.SetMessage(queue.ID, *message)
	mbs.CallSubscribers(queue, message)
	return nil, nil
}
func (mbs *MessageBrokerService) GetMessage(queue *domain.Queue, subscriberID string, messageID int) domain.Message {
	logger.Logger.Debug().Msg("Message Broker GetMessage")
	msg := mbs.MessageStoreService.GetMessage(messageID)
	return msg
}
func (mbs *MessageBrokerService) DeleteMessage() {}
func (mbs *MessageBrokerService) CallSubscribers(queueRef *domain.Queue, message *domain.Message) error {
	logger.Logger.Debug().Msg("Message Broker callsubscribers")
	mbs.CallBackChan <- message
	return nil
}
