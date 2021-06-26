package service

type ISenderService interface {
	GetMessage()
	RegisterSubscriber()
}
