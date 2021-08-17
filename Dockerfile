FROM golang:1.16-alpine AS builder
RUN go version

COPY . /app
WORKDIR /app
RUN set -x && \
    go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]