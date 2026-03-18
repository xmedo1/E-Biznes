FROM ubuntu:24.04
ENV TZ=Europe/Warsaw
WORKDIR /app

RUN apt-get update && apt-get install -y openjdk-8-jdk kotlin

ENV JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64

CMD ["/bin/bash", "-c", "java -version && kotlinc -version"]