package service

import "nipun.io/message_queue/domain"

type MessageStoreService struct {
	queueToMessageListMap map[string][]domain.Message
}

func (mss *MessageStoreService) GetMessage() {}
func (mss *MessageStoreService) SetMessage(queueName string, message domain.Message) error {
	return nil
}
func (mss *MessageStoreService) DeleteMessage() {}
