FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
COPY . .
ENTRYPOINT ["air"]