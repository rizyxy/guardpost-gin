FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git gcc musl-dev # GCC is needed for SQLite drivers
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o guardpost .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

RUN mkdir /root/data

COPY --from=builder /app/guardpost .
COPY --from=builder /app/routes.yaml .
COPY --from=builder /app/.env . 

EXPOSE 8080

CMD ["./guardpost"]