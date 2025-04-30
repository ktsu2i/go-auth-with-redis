FROM golang:1.23-bookworm

WORKDIR /backend

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum /backend/
RUN go mod download

COPY . /backend
COPY .air.toml /backend/.air.toml

CMD ["sh", "-c", "air -c .air.toml"]
