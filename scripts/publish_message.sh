#!/bin/bash

curl --location --request POST 'http://localhost:8080/v1/publish' \
--header 'Content-Type: application/json' \
--data-raw '{
    "metadata" : {
        "queue_name":"my-queue"
    },
    "value" : {
        "qwerty" : 123,
        "qpwoei": "qwerty"
        }
}'