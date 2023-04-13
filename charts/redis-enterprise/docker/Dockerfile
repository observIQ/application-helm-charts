FROM golang:latest as go

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /gocurl

FROM redislabs/redis:6.4.2-61.focal

USER root

COPY --from=go /gocurl /gocurl
RUN chmod +x /gocurl

RUN apt-get update
RUN apt-get upgrade curl -y

USER redislabs
