#!/usr/bin/bash
docker build -t zadanie08-ebiznes .
docker run -it --rm -p 8080:8080 -p 5173:5173 --env-file .env zadanie08-ebiznes