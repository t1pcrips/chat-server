FROM golang:1.23.3-alpine AS builder

COPY . /github.com/t1pcrips/chat-server/source/
WORKDIR /github.com/t1pcrips/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/t1pcrips/chat-server/source/bin/crud_server .

CMD ["./chat_server"]