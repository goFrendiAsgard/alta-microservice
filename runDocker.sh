#!/bin/bash

# FASE BUILD
docker build -t app .


# FASE DEPLOY
# remove old container
docker stop appReader
docker rm appReader
docker stop appWriter
docker rm appWriter
docker stop appGateway
docker rm appGateway

# create and run new container
docker run -d --name appReader -e MODE=reader -e READER_PORT=3000 -p 3000:3000 app
docker run -d --name appWriter -e MODE=writer -e WRITER_PORT=5000 -p 5000:5000 app
docker run -d --name appGateway -e MODE=gateway -e GATEWAY_PORT=8080 -e READER_URL=http://host.docker.internal:3000 -e WRITER_URL=http://host.docker.internal:5000 -p 8080:8080 app