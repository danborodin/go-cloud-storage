FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./bin/auth ./src/cmd/main.go
EXPOSE 6969
CMD ["./bin/auth", "-l=/var/log/auth.log"]
#CMD ["./bin/auth"]