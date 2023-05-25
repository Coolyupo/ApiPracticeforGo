FROM golang:1.20.4-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GIN_MODE=release
RUN CGO_ENABLED=0 GOOS=linux go build -o /myapi
CMD ["/myapi"]