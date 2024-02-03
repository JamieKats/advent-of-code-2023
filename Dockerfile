FROM golang:alpine3.19 as base

FROM base as dev

RUN go install github.com/cosmtrek/air@latest

WORKDIR /opt/app/api

RUN air init

# CMD [ "air" ]