FROM golang:1.24.2-alpine3.21

WORKDIR /app/backend

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum /app/backend/
RUN go mod download

COPY . /app/backend
COPY .air.toml /app/backend/.air.toml

CMD ["sh", "-c", "air -c .air.toml"]
