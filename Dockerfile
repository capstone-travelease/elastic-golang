FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o elastic-golang

# This one will reduce size of project image
FROM scratch

COPY --from=builder /app/elastic-golang /

ENTRYPOINT ["/elastic-golang"]
