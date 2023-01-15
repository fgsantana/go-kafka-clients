FROM golang:1.19

WORKDIR /app
COPY main.go .
COPY go.mod .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o /consumer main.go

CMD [ "/consumer" ]