FROM golang:1.23 AS base
WORKDIR /app

# Run in development
FROM base AS dev
RUN go install github.com/air-verse/air@latest
COPY go.* ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["air"]

# Run in production
FROM base AS prod
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o 1up
EXPOSE 8080
CMD ["/app/1up"]
