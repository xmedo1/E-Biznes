#!/usr/bin/bash
docker build -t zadanie04-ebiznes .
docker run --rm -p 8080:8080 -v "$(pwd)/database.db:/app/database.db" zadanie04-ebiznes