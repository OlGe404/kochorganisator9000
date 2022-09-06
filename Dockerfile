FROM docker.io/library/golang:1.19

WORKDIR /usr/src/app

COPY ./app/ ./

RUN go mod download && \
    go mod verify && \
    go build -v -o /usr/local/bin/app ./

CMD ["app"]
