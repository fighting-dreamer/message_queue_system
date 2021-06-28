#!/bin/bash

curl --location --request POST 'http://localhost:8080/v1/subscribe/poll' \
--header 'Content-Type: application/json' \
--data-raw '{
    "subscriber_id":"sub-1",
    "message_id":2,
    "fetch_count":5
}'