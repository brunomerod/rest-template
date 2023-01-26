#!/bin/bash

curl --location --request POST 'localhost:8083/transactions' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id":"629cfe79b9afb07e622189fb",
    "operation_type_id":1,
    "amount":122223.45
}'