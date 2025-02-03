FROM golang:1.23.5-alpine
WORKDIR /go/src/app
COPY . .
CMD ["go", "run", "cmd/rest/main.go"]