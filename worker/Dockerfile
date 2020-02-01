FROM golang

ENV RABBITMQ_URI amqp://guest:guest@rabbitmq:5672
ENV PORT 2500

ADD . /go/src/worker

WORKDIR /go/src/worker

RUN go install

ENTRYPOINT ["/go/bin/worker"]