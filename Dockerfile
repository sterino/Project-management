FROM golang:1.22.5-alpine AS builder

WORKDIR /build

#COPY go.mod go.sum ./
#
#RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o project-management ./cmd


FROM alpine AS hoster

COPY --from=builder /build/.env .env
COPY --from=builder /build/project-management ./project-management
COPY --from=builder /build/migrations ./migrations

ENTRYPOINT ["./project-management"]
