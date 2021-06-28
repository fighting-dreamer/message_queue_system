package service

import (
	"errors"
	"nipun.io/message_queue/domain"
)

var (
	MessageDoesNotExist = errors.New("MessageDoesNotExist")
)

type MessageStoreService struct {
	MessageIDMap            map[int64]domain.Message
	QueueToMessageIDListMap map[string][]int64
	QueueCounter            map[string]int64
}

func (mss *MessageStoreService) GetMessage(messageID int64) (domain.Message, error) {
	// TODO : proper algorithm to be implemented
	message := mss.MessageIDMap[messageID]
	if message.ID == messageID {
		return message, nil
	}
	return domain.Message{}, MessageDoesNotExist
}

func (mss *MessageStoreService) SetMessage(queueName string, message domain.Message) (domain.Message, error) {
	// TODO : acquire lock on QueueCounter
	// TODO : acquire lock on MessageIDMap
	// TODO : acquire lock on QueueToMessageIDListMap
	message.ID = mss.QueueCounter[queueName] + 1
	mss.MessageIDMap[message.ID] = message
	mss.QueueToMessageIDListMap[queueName] = append(mss.QueueToMessageIDListMap[queueName], message.ID)
	return message, nil
}
