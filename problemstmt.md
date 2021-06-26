## Developing a message queueing system.
Functional requirements of this system have been described below.
1. Create your own queue that will hold messages in the form of JSON(Standard
library with queue implementation were not allowed). 
2. There can be more than one queue at any given point of time.
3. There will be one publisher that can generate messages to a queue. 
4. There are multiple subscribers that will listen to queues for messages.
5. Subscribers should not be tightly coupled to the system and can be added or
removed at runtime.
6. When a subscriber is added to the system, It registers a callback function
which makes an API call to the given end point with the json payload, this
callback function will be invoked in case some message arrives.
7. Subscriber can consume the messages in batches if the queue has more than
one message and it should be configurable. 
> Bonus: 
There must be a retry mechanism for handling error cases when some
exception occurs in listening/processing a message, that must be retried.

## NOTE : Things to keep in mind while submitting the code: 
1. Make sure that you can execute the code and show that in a working state
2. Use an in-memory store for now. No need to use any database to store data
3. Make sure code functionality is correct, and edge cases are covered
4. Code should be easily testable
5. Exceptions are raised and handled appropriately.
6. Separation of concerns is addressed.