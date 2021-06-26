package service

type IMessageStoreService interface {
	GetMessage()
	SetMessage()
	DeleteMessage()
}
