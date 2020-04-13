FROM ubuntu:18.04

WORKDIR /bin

ADD golang-pipeline /bin/golang-pipeline

RUN chmod +x /bin/golang-pipeline

ENTRYPOINT /bin/golang-pipeline

EXPOSE 8080
