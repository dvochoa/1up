FROM golang:1.23

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o 1up

EXPOSE 8080

CMD ["/app/1up"]
