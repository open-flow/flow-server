FROM golang:1.16-alpine AS builder

COPY . /app
WORKDIR /app
RUN set -x && \
    go mod tidy

RUN go get -u github.com/swaggo/swag/cmd/swag  && \
    swag init -g internal/http/controller.go


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]
CMD ["serve"]