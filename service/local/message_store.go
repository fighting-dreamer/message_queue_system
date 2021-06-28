package service

import (
	"errors"
	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

var (
	MessageDoesNotExist = errors.New("MessageDoesNotExist")
)

type MessageStoreService struct {
	MessageIDMap            map[int64]domain.Message
	QueueToMessageIDListMap map[string][]int64
	QueueCounter            map[string]int64
	TransactionLockManager  service.ITransactionLockManager
}

func (mss *MessageStoreService) GetMessage(messageID int64) (domain.Message, error) {
	// TODO : proper algorithm to be implemented
	mss.TransactionLockManager.AcquireLock([]string{"MessageIDMap"})

	message := mss.MessageIDMap[messageID]

	mss.TransactionLockManager.ReleaseLock([]string{"MessageIDMap"})

	if message.ID == messageID {
		return message, nil
	}
	return domain.Message{}, MessageDoesNotExist
}

func (mss *MessageStoreService) SetMessage(queueName string, message domain.Message) (domain.Message, error) {
	// TODO : acquire lock on QueueCounter
	// TODO : acquire lock on MessageIDMap
	// TODO : acquire lock on QueueToMessageIDListMap
	mss.TransactionLockManager.AcquireLock([]string{"QueueCounter", "MessageIDMap", "QueueToMessageIDListMap"})
	logger.Logger.Debug().Msgf("Setting started %+v : ", message)
	message.ID = mss.QueueCounter[queueName] + 1
	mss.QueueCounter[queueName] = message.ID
	mss.MessageIDMap[message.ID] = message
	mss.QueueToMessageIDListMap[queueName] = append(mss.QueueToMessageIDListMap[queueName], message.ID)
	logger.Logger.Debug().Msgf("Setting completed %+v : ", message)
	mss.TransactionLockManager.ReleaseLock([]string{"QueueCounter", "MessageIDMap", "QueueToMessageIDListMap"})
	return message, nil
}
