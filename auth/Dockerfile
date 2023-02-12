FROM golang:1.20
WORKDIR /app
COPY src ./src
COPY go.mod go.mod
COPY go.sum go.sum
COPY dev.env .env
RUN go mod download
RUN go build -o ./bin/auth ./src/cmd/main.go
EXPOSE 6969
CMD ["./bin/auth", "-l=/var/log/auth.log"]
#CMD ["./bin/auth"]