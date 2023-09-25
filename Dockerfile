FROM golang:1.21.1-alpine

WORKDIR /app

# Download the go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /sweeTooth

EXPOSE 8080

CMD ["/sweeTooth"]

