FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o billing cmd/billing/main.go

FROM scratch as run
COPY --from=builder /app/billing /
ENTRYPOINT ["/billing"]
