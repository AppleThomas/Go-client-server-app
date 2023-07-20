FROM golang:1.19
WORKDIR /app

# Download Go Modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY *.go ./

RUN go build -o /go-client-server

EXPOSE 3000

CMD ["/go-client-server"]