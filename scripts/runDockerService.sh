#!/usr/bin/env bash
docker run -d -p 8090:30200 -e AUTHENTICATION_AUDIENCE='fir-sample-9b54d' game_service:latest


