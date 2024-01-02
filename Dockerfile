FROM golang:1.21

WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go install github.com/cosmtrek/air@latest

COPY .env .env

EXPOSE 2160
CMD ["air"]