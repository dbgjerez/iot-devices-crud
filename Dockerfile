FROM golang:alpine

ENV APP_NAME=ms-iot-devices-crud
ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE $PORT

ENTRYPOINT $APP_NAME
