#!/bin/bash

curl --location --request POST 'http://localhost:8080/v1/subscribe/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "queue_name": "my-queue",
    "callback_url" : "http://localhost:3003/msg",
    "subscriber_id": "sub-1"
}'