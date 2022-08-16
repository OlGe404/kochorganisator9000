FROM golang:1.19

WORKDIR /usr/src/app

COPY hello_world.go ./

RUN go build -v -o /usr/local/bin/app ./

CMD ["app"]
