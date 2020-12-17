FROM golang:1.15

WORKDIR /go/src/app
COPY . .

ENV GIT_TERMINAL_PROMPT=1
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]