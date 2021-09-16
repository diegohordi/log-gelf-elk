FROM golang:1.16.7-alpine3.14 as build
ENV GOOS linux
ENV CGO_ENABLED 0
RUN mkdir /app
COPY /cmd/restapi/main.go /app/cmd/restapi/main.go
COPY go.mod /app
WORKDIR /app
RUN go mod download
RUN go build -o restapi cmd/restapi/main.go


FROM alpine:3.14 as deploy
ARG SERVER_PORT
ENV SERVER_PORT=$SERVER_PORT

RUN apk add --no-cache curl
RUN apk add --no-cache ca-certificates
COPY --from=build /app/restapi /app/restapi

EXPOSE $SERVER_PORT
CMD cd /app/ && ./restapi
