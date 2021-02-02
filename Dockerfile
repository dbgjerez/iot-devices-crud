FROM golang:1.15-alpine

ENV APP_NAME=iot-devices-crud
ENV GIN_MODE=release
ENV PORT=8080
ENV MONGODB_DEVICE_DB=iot
ENV MONGODB_DEVICE_COLLECTION_NAME=device

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN rm -rf /go/src/app

EXPOSE $PORT

ENTRYPOINT $APP_NAME
