FROM golang:1.22.5-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o project ./cmd

EXPOSE 8080

CMD ["./project"]

FROM alpine AS hoster

WORKDIR /app

COPY --from=builder /build/migrations ./migrations
COPY --from=builder /build/project ./project

EXPOSE 8080

CMD ["./project"]
