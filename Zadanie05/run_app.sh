#!/usr/bin/bash
docker build -t zadanie05-ebiznes .
docker run -it --rm -p 8080:8080 -p 5173:5173 zadanie05-ebiznes