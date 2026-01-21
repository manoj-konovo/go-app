FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o go-app ./main.go

FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/go-app /usr/local/bin/go-app

EXPOSE 8080
ENV PORT=8080

USER nobody

ENTRYPOINT ["go-app"]