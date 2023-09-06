FROM golang:1.20 as builder

LABEL maintainer="alex.vidyakov@yandex.ru"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/gophermart
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.18.3 as production

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/gophermart/main ./

EXPOSE 8080

CMD ["./main"]
