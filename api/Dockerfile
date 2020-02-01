FROM golang

ENV RABBITMQ_URI amqp://guest:guest@rabbitmq:5672
ENV PORT 2500

ADD . /go/src/api

WORKDIR /go/src/api

RUN go install

ENTRYPOINT ["/go/bin/api"]