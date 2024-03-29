package appcontext

import (
	resty "github.com/go-resty/resty/v2"
	"nipun.io/message_queue/domain"
	service "nipun.io/message_queue/service"
	local_service "nipun.io/message_queue/service/local"
	"sync"
)

type Instance struct {
	CallBackChan             chan *domain.Message
	MessageStoreService      service.IMessageStoreService
	QueueManager             service.IQueueManager
	MessageBroker            service.IMessageBrokerService
	RecieverService          service.IRecieverService
	SubscriberManager        service.ISubscriberManager
	SenderService            service.ISenderService
	CallbackWorker           service.ICallBackWorker
	TransactionalLockManager service.ITransactionLockManager
}

var AppDependencies *Instance

func LoadDependencies() {
	AppDependencies = &Instance{
		CallBackChan: make(chan *domain.Message),
	}
	addTransactionalLockManager(AppDependencies)
	addMessageStore(AppDependencies)
	addQueueManager(AppDependencies)
	addMessageBroker(AppDependencies)
	addRecieverService(AppDependencies)
	addSubscriberManager(AppDependencies)
	addSenderService(AppDependencies)
	addCallbackWorker(AppDependencies)
}

func addTransactionalLockManager(dependencies *Instance) {
	dependencies.TransactionalLockManager = &local_service.TransactionLockManager{
		KeeperState: map[string]string{},
		Keeper:      map[string]*sync.Mutex{},
	}
}

func addMessageStore(dependencies *Instance) {
	dependencies.MessageStoreService = &local_service.MessageStoreService{
		MessageIDMap:            map[int64]domain.Message{},
		QueueToMessageIDListMap: map[string][]int64{},
		QueueCounter:            map[string]int64{},
		TransactionLockManager:  dependencies.TransactionalLockManager,
	}
}

func addQueueManager(dependencies *Instance) {
	queueMap := map[string]*domain.Queue{}
	dependencies.QueueManager = &local_service.QueueManager{
		QueueMap:               queueMap,
		TransactionLockManager: dependencies.TransactionalLockManager,
	}
}

func addMessageBroker(dependencies *Instance) {
	dependencies.MessageBroker = &local_service.MessageBrokerService{
		MessageStoreService: dependencies.MessageStoreService,
		CallBackChan:        dependencies.CallBackChan,
	}
}

func addRecieverService(dependencies *Instance) {
	dependencies.RecieverService = &local_service.RecieverService{
		QueueManager:  dependencies.QueueManager,
		MessageBroker: dependencies.MessageBroker,
	}
}

func addSubscriberManager(dependencies *Instance) {
	dependencies.SubscriberManager = &local_service.SubscriberManager{
		UnackCounterPerQueuePerSubscriber: map[string]int64{},
		AckCounterPerQueuePerSubscriber:   map[string]int64{},
		QueueManager:                      dependencies.QueueManager,
		SubscriberMap:                     map[string]*domain.Subscriber{},
		SubscriberToQueueMap:              map[string]string{},
		QueueToSubscriberListMap:          map[string][]*domain.Subscriber{},
		TransactionLockManager:            dependencies.TransactionalLockManager,
	}
}

func addSenderService(dependencies *Instance) {
	dependencies.SenderService = &local_service.SenderService{
		QueueManager:      dependencies.QueueManager,
		SubscriberManager: dependencies.SubscriberManager,
		MessageBroker:     dependencies.MessageBroker,
	}
}

func addCallbackWorker(dependencies *Instance) {
	dependencies.CallbackWorker = &local_service.CallBackWorker{
		SubscriberManager: dependencies.SubscriberManager,
		CallBackChan:      dependencies.CallBackChan,
		HttpCli:           resty.New(),
	}
}
