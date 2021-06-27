package service

import "nipun.io/message_queue/domain"

type IRecieverService interface {
	EnqueueMessage(message domain.Message)
}
