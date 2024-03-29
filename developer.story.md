# Algorithm

##NOTE : 

- I may have over-engineered this whole problem. 
- I have tried to ensure a concurrent publish, concurrent subscriber.
- I have implemented few things inspired from rabbitMQ, kafka and TCP protocol. 
- Some Checkpoints using git tags.

### Objective :
1. we had options of Blocking queue on subsciber fetch OR not-blocking queue on subscriber's fetch, we have tried to implement the **non-blocking queue**
2. we had the option of synschronous processing of messages by subscribers OR async processing and acks by subscribers, we have tried to give **async prcessing and acks** by subscriber.

### Concerns and Assumptions :
- publisher api must ensure some kind of rate limiting to hold off the publisher from bombarding messages.
- publisher is very simple client, just send the message.
- subscriber can get the message itself in the via the call-back api.
- subscriber will and can take its own time to consumer messages.
- subscriber can go down as it is processing and want to re-try.
- subscriber want to start from index it left off.
- subscriber want to consumer at its own pace.
- we want to hld off sending new messages if the subscriber is not acknoledging the messages and thus keep the processing of new messages in check.
- subscriber is intelligent enough the it know if it have processed certain message or not but keep track of this info to a given time-period only.
- we want to follow "at-least once" delivery sementics.
- subscriber starts to poll via the call-back api initiation.
- subscriber should be able to poll a particular message as a re-try operation.
- broker should ensure quick relay of messages as they get published.
- broker may not need to deal with any kind of batching in call-back api for a given subscriber.
- broker don;t bother with response from the call-back api.
- broker will want to send hte messages and don;t bother with their processing in batches
- broker service will get acks async.
- we are not providing the delete queue kind of api, but the system should be able to handle that properly.
- From publisher perspective our queue is going to be blocking as it does not let one api call to happen in parallel to other for publish operation.
- subscriber will be accessing the queue data concurrently.
- In this IF we go with retry at the server/broker level,
    - it will incur re-processing cost to all the subscribers if we put the message back in queue.
    - it will require to re-process form the point where the message was not ack.
    - it will require to copy messages across subscribers and keep their track seperately.
    - it will require to keep a re-try queue peer subscriber for a given main queue.

### [Rough] : thoughtprocess :
1. For publisher :
- there could be many publishers tomo, we want to just publish the message and get success.
- the queue is a shared resource, so modification in that will require locks.

2. For Subscriber :
- If we want to just send the data to a subscriber then we can send the enqueueed message directly to the subscriber.
    - It will work even if there are many subscribers registered for that queue at that time.
- subscriber gets a callback just as a notification that a message was added, it may not be able to process the message right away.
- If the subscriber is going to poll then it implies, we are trying to mitigate away from the synchronous processing of the message due to callback api.
- Scenarios :
    - Given :
        - the publisher rate and subscriber's rate of processing can differ a lot.
        - the subscriber is going to poll and it can get a batch of messages too.
        - on top of this, the poll is initiated after the call-back api.
    - IF publisher rate > subscriber rate
        - the queue length will increase a lot.
        - the subscriber will ignore/have to ignore the call-back api calls as it will get a lot of such calls.
        - Action : we want to cap the max messages in a queue that are ack and un-ack.
    - IF pubclisher rate <= subscriber rate
        - the queue length will be essentially be consumed up.
        - most poll batch request need to send only the left out message.

### core compnent : MessageBroker :
1. SetMessage :
    1. acquire a lock on queue
    2. create a incremental message ID as (queue-length + 1)
    3. add the message in respective place. // TODO : what should be those respective places.
    4. release the lock on queue.
    5. Async call to call_worker for communicating the subscribers.
    - Communicate what ? Options:
        1. Latest message ID.
        2. latest message itself.
        3. some index that help the subscriber know till what the message have been consumed OR should start
            1. communicate the min of index of latest un-ack message. **[CHOOSEN]**
    6. return with message with an ID.


2. GetMessage :
    1. acquire a lock on subscriberID so that it can be operated on, without any issue of other similar api calls from same subscriber.
    2. read the messages from the queue. HOW many ?
        1. the subscriber wants 'x' messages from index i, so **messages[i : i + x]** .
        2. the un-ack counter tells what messages were already sent out to subscriber, it may coincide with the message[i : i + x] range.
        3. so we want to send out messages range **(max(un-ack counter, i), i + x]** (this is ordered set, with un-ack counter message as open)
        4. sending this many message will update the un-ack counter to i + x **given the i + x > un-ack counter** .
        5. with that the difference b/w the ack-counter and un-ack counter may increses than the **queue cap** defined for a given subscriber for that queue of total un-ack message
        6. then we will need to update the range to **(max(un-ack counter, i), min(ack counter + cap_per_subscriber_per_queue, i + x)]** .
    3. Operate the queue to get the messages and update the counters. The un-ack counter and ack counters need to be under lock too.
    4. send the messages back.

### Still problems :
1. The problem of concurrent get with this strategy is still tricky, one way is to have more intelligence built in subscriber OR have coordination taken seperately rather then here in broker.
2. we had options of Blocking queue on subsciber fetch OR not-blocking queue on subscriber's fetch, we have tried to implement the **non-blocking queue**
3. we had the option of synschronous processing of messages by subscribers OR async processing and acks by subscribers, we have tried to give **async prcessing and acks** by subscriber.
