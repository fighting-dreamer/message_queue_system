package service

type MessageBrokerService struct {
}

func (mbs *MessageBrokerService) SetMessage()      {}
func (mbs *MessageBrokerService) GetMessage()      {}
func (mbs *MessageBrokerService) DeleteMessage()   {}
func (mbs *MessageBrokerService) CallSubscribers() {}
