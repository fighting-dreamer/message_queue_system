package service

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"

	"nipun.io/message_queue/domain"
	"nipun.io/message_queue/logger"
	"nipun.io/message_queue/service"
)

type CallBackWorker struct {
	SubscriberManager service.ISubscriberManager
	HttpCli           *resty.Client
	CallBackChan      chan *domain.Message
}

func Start(worker service.ICallBackWorker) {
	logger.Logger.Debug().Msg("Started the Callback Worker")
	for messageRef := range worker.GetCallBackChan() {
		logger.Logger.Debug().Msgf("Got Message : %+v", *messageRef)
		go worker.CallSubscribers(messageRef)
	}
}

func (cw *CallBackWorker) CallSubscribers(message *domain.Message) error {
	queueName := message.Metadata.QueueName
	subscribers := cw.SubscriberManager.GetQueueSubscribers(queueName)
	for _, subscriber := range subscribers {
		logger.Logger.Debug().Msgf("Trying to message : %+v, for subsciber : %+v", message, subscriber)
		go func() {
			url := subscriber.URL
			res, err := json.Marshal(message)
			if err != nil {
				// log and increment and return
				logger.Logger.Debug().Msgf("Got Error %s", err.Error())
				return
			}

			response, err := cw.HttpCli.R().
				SetHeader("Content-Type", "application/json").
				SetBody(res).
				Post(url)

			if err != nil {
				// Do something
				logger.Logger.Debug().Msgf("Got Response Error %s", err.Error())
			}
			logger.Logger.Debug().Msgf("Got Response : %s", response.Status())
			logger.Logger.Debug().Msgf("Got Response : %s", response.Body())

			// TODO : use response status to update metrics and log it
			// TODO : ignore response body for now.
		}()
	}

	return nil
}

func (cw *CallBackWorker) GetCallBackChan() chan *domain.Message {
	return cw.CallBackChan
}
