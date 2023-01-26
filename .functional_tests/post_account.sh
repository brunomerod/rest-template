#!/bin/bash

curl --location --request POST 'localhost:8083/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "document_number":"12334564400"
}'