# https://stackoverflow.com/questions/47837149/build-docker-with-go-app-cannot-find-package
FROM golang:alpine

ENV APP_NAME banking

WORKDIR /go/src/${APP_NAME}
COPY . .
WORKDIR /go/src/${APP_NAME}/cmd
RUN go build -o ${APP_NAME}
ENTRYPOINT ["./banking"]