FROM ubuntu:24.04
ENV TZ=Europe/Warsaw

RUN apt-get update && apt-get install -y apt-transport-https curl gnupg openjdk-17-jdk

RUN curl -fL https://github.com/sbt/sbt/releases/download/v1.12.8/sbt-1.12.8.tgz | tar -xz -C /usr/local
ENV PATH="/usr/local/sbt/bin:$PATH"

WORKDIR /app

COPY ./my-scalatra-web-app /app

RUN sbt compile
EXPOSE 8080

CMD ["sbt", "~warStart"]