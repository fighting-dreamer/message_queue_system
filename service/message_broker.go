package service

type IMessageBrokerService interface {
	SetMessage()
	GetMessage()
	DeleteMessage()
	CallSubscribers()
}
