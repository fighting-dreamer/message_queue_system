## Message Queue System

## Resources : 
1. Queue
2. Message
3. Publisher
4. Subscriber
5. Server

## Services :
1. MessageStore
2. QueueManager
3. QueueHandler
4. MessageBroker
5. RecieverService
6. PublisherHandler
7. SenderService
8. SubscriberManager
9. SubscriberHandler

## Other Aspects : 
1. Logging
2. HttpServerMux
3. Monitoring
4. Security
5. Exception Handling
6. Utils
7. Config Management
8. Dependency Management

## Other related project : 
1. publisher client.
2. subscriber application.

### Details :
1. MessageStore
    1. GetMessage()
    2. SetMessage()
    3. DeleteMessage()

2. QueueManager
    1. CreateQueue()
    2. GetQueue()

3. QueueHandler 
    - QueueManager
    1. CreateQueueApi()

4. MessageBroker
    1. SetMessage()
    2. GetMessage()
    3. DeleteMessage()
    4. CallSubscribers()

5. RecieverService
    - QueueManager
    - MessageBroker
    1. EnqueueMessage()

6. PublisherHandler
    - RecieverService
    1. PublishMessageApi()

7. SubscriberManager
    1. RegisterSubscriber()
    3. GetQueueSubscribers()

8. SenderService
    - QueueManager
    - MessageBroker
    1. GetMessage()

9. SubscriberHandler
    - SenderService
    1. RegisterSubscriberApi()
    2. PollMessageApi()

## Processes :
1. CreateQueue
2. RegisterSubscriber
3. PublishMessage
4. PollMessage

-[x] CreateQueue
    1. User calls QueueHandler.**CreateQueueApi()**.
    2. It calls QueueManager.**CreateQueue()** to creates a queue
    3. CreateQueue checks if the queue already existed ?
       1. IF yes then raise Exception
       2. Else return success.
    4. Pass the result to handler to create response.

-[x] RegisterSubscriber
   1. User calls SubscriberHandler.**RegisterSubscriberApi()**.
   2. It calls the subscriberManager.**RegisterSubscriber()**.
   3. It checks if the Queue exists using QueueManager?
      1. IF yes then continue
      2. Else raise exception and return.
   4. It checks if the subsciber is already registered for that queue using SubscriberManager?
      1. If yes then raise exception and return
      2. Else add subsciber in it and return success.

-[x] PublishMessage
   1. User publish message using curl/http api on PublisherHandler.**PublishMessageApi()**.
   2. It calls the RecieverService.**EnqueueMessage()**.
   3. It validates the message queue using QueueManager?
      1. IF Queue exists then continue
      2. Else raise exception and return.
   4. It calls MessageBroker.**SetMessage()** and return success OR exception.
   5. If the publish is successful then it raises an async operation on the MessageBroker.**CallSubscribers()**
   
4. PollMessage
   1. Subscriber calls SubscriberHandler.**PollMessageApi()**.
   2. It calls then SenderService.GetMessage()
   3. It checks the queue is valid. 
      1. IF yes continue, 
      2. Else raise exception and return
   4. It checks the subscriber is registered for that queue or not ?
      1. If yes then continue
      2. Else raise exception and return
   5. It calls MessageBroker.**GetMessage()** to get the messages

