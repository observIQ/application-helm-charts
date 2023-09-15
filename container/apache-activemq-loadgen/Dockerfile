FROM golang AS build

WORKDIR /app

COPY go.mod go.sum receiver-producer.go ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /receiver-producer

FROM alpine:latest
WORKDIR /
COPY --from=build /receiver-producer /receiver-producer

ENTRYPOINT [ "/receiver-producer" ]
