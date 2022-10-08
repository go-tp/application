#!/bin/bash

docker stop go
docker rm go
docker run --name go -d  -v /www/go/:/go/src  -p 778:7777 -it go/1.19