#!/bin/bash

curl --location --request POST 'http://localhost:8080/v1/subscribe/message/ack' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message_id":1,
    "queue_name":"my-queue",
    "subscriber_id":"sub-1"
}'