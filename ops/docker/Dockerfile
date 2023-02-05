FROM golang:1.19

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

VOLUME /app

RUN go build -o ./go-cli

ENTRYPOINT ["./go-cli"]