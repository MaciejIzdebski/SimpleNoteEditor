# syntax=docker/dockerfile:1.2
FROM golang:1.24.2-alpine AS build

WORKDIR /go/src/app

COPY . .


RUN apk add build-base

ENV CGO_ENABLED=1
RUN go build -o /bin/note_server ./main.go

FROM alpine


COPY --from=build /bin/note_server /bin/note_server

EXPOSE 9090

ENTRYPOINT [ "/bin/note_server" ]