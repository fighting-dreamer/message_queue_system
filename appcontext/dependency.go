package appcontext

import (
	service "nipun.io/message_queue/service"
	local_service "nipun.io/message_queue/service/local"
)

type Instance struct {
	MessageStore      service.IMessageStoreService
	QueueManager      service.IQueueManager
	MessageBroker     service.IMessageBrokerService
	RecieverService   service.IRecieverService
	SubscriberManager service.ISubscriberManager
	SenderService     service.ISenderService
}

var AppDependencies *Instance

func LoadDependencies() {
	addMessageStore(AppDependencies)
	addQueueManager(AppDependencies)
	addMessageBroker(AppDependencies)
	addRecieverService(AppDependencies)
	addSubscriberManager(AppDependencies)
	addSenderService(AppDependencies)
}

func addMessageStore(dependencies *Instance) {
	dependencies.MessageStore = &local_service.MessageStoreService{}
}

func addQueueManager(dependencies *Instance) {
	dependencies.QueueManager = &local_service.QueueManager{}
}

func addMessageBroker(dependencies *Instance) {
	dependencies.MessageBroker = &local_service.MessageBrokerService{}
}

func addRecieverService(dependencies *Instance) {
	dependencies.RecieverService = &local_service.RecieverService{
		QueueManager:  dependencies.QueueManager,
		MessageBroker: dependencies.MessageBroker,
	}
}

func addSubscriberManager(dependencies *Instance) {
	dependencies.SubscriberManager = &local_service.SubscriberManager{}
}

func addSenderService(dependencies *Instance) {
	dependencies.SenderService = &local_service.SenderService{
		QueueManager:      dependencies.QueueManager,
		SubscriberManager: dependencies.SubscriberManager,
		MessageBroker:     dependencies.MessageBroker,
	}
}
