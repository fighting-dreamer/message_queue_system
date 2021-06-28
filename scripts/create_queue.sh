#!/bin/bash

curl --location --request POST 'http://localhost:8080/v1/queue/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "my-queue"
}'