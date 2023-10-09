FROM docker.io/library/golang:1.21.1-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM gcr.io/distroless/static as release
WORKDIR /app
COPY --from=builder /app .

ENTRYPOINT ["./app"]

