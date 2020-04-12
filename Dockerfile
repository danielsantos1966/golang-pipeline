FROM alpine

WORKDIR /bin

ADD golang-pipeline /bin/golang-pipeline
ADD . /go/src/golang-pipeline

RUN chmod +x /bin/golang-pipeline

ENTRYPOINT /bin/golang-pipeline

EXPOSE 8080
