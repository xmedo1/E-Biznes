#!/usr/bin/bash
docker build -t my-scalatra-web-app .
docker run -d -p 8080:8080 my-scalatra-web-app
ngrok http 127.0.0.1:8080