ARG BASE_IMAGE=alpine:3.20

FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app

RUN apk add upx

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o dankey
RUN upx dankey

FROM $BASE_IMAGE

WORKDIR /app
COPY --from=builder /app/dankey .
EXPOSE 6969

CMD ["./dankey"]
