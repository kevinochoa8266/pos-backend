FROM golang:1.21.2-alpine as build

WORKDIR /src

# Download the go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -o sweeTooth .

FROM alpine:latest
RUN mkdir /data
COPY --from=build /src/sweeTooth /sweeTooth
COPY --from=build  /src/.env .
COPY --from=build /src/candy_data.csv .
EXPOSE 8080
CMD ["/sweeTooth"]

