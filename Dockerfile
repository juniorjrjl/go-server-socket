FROM golang:1.22.5-bullseye

RUN apt-get update && apt-get install -qq -y --no-install-recommends

RUN go install github.com/jackc/tern/v2@v2.2.1
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0

ENV INSTALL_PATH /go-server-socket

RUN mkdir $INSTALL_PATH

WORKDIR $INSTALL_PATH

COPY . .