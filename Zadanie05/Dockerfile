FROM debian:stable-slim

RUN apt-get update && apt-get install -y nodejs npm golang-go

WORKDIR /app

COPY sklep-backend ./sklep-backend
RUN cd sklep-backend && go build -o server .

COPY sklep ./sklep
RUN cd sklep && npm install

EXPOSE 8080 5173

CMD sh -c "/app/sklep-backend/server & cd /app/sklep && npm run dev -- --host"