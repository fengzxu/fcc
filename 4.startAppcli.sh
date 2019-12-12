#!/bin/sh
cd appcode
docker-compose -f docker-compose-appcli.yaml up -d
cd ..