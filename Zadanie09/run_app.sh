#!/usr/bin/bash
python PythonApp/app.py &
sleep 2

docker build -t zadanie09-ebiznes .
docker run --rm --network host --env-file DiscordBotApp/.env zadanie09-ebiznes