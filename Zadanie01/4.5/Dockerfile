FROM ubuntu:24.04
ENV TZ=Europe/Warsaw
WORKDIR /app

RUN apt-get update && apt-get install -y openjdk-17-jdk kotlin wget unzip

ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64

RUN wget https://services.gradle.org/distributions/gradle-9.4.0-bin.zip -P /tmp && unzip -d /opt/gradle /tmp/gradle-9.4.0-bin.zip && rm /tmp/gradle-9.4.0-bin.zip
ENV PATH=$PATH:/opt/gradle/gradle-9.4.0/bin

COPY build.gradle .
COPY src ./src

CMD ["gradle", "run"]