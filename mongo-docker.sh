#!/bin/bash
# Para facilitar o uso local do MongoDB

docker run -d -e MONGO_INITDB_ROOT_USERNAME=lmongo -e MONGO_INITDB_ROOT_PASSWORD=lmongo -p 27017:27017 mongo