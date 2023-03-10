FROM golang:1.20-alpine as compile
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o payment cmd/payment/main.go

FROM scratch
COPY --from=builder /app/payment /
ENTRYPOINT ["/payment"]
