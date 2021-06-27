package service

import (
	"math/rand"
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
)

type MessageStoreService struct {
	MessageIDMap            map[int]domain.Message
	QueueToMessageIDListMap map[string][]int
}

func (mss *MessageStoreService) GetMessage(messageID int) domain.Message {
	// TODO : proper algorithm to be implemented
	var tempMsg domain.Message
	flag := false
	for k, v := range mss.MessageIDMap {
		logger.Logger.Debug().Msgf("Key : %d, Value : %+v", k, v)
		if rand.Int31n(1000) > 500 && flag == false {
			tempMsg = v
			flag = true
		}
	}
	return tempMsg
}
func (mss *MessageStoreService) SetMessage(queueName string, message domain.Message) error {
	message.ID = int(rand.Int31n(1000))
	mss.MessageIDMap[message.ID] = message
	mss.QueueToMessageIDListMap[queueName] = append(mss.QueueToMessageIDListMap[queueName], message.ID)
	return nil
}
func (mss *MessageStoreService) DeleteMessage() {}
