#!/usr/bin/bash
docker build -t zadanie03-ebiznes .
docker run --rm -p 8080:8080 --env-file DiscordBotApp/.env zadanie03-ebiznes