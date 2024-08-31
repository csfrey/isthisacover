FROM golang:1.22.6-alpine3.19

WORKDIR /app

COPY . .

# download and install the dependencies
RUN go get -d -v ./...

# build the go app
RUN go build -o api .

EXPOSE 8000

CMD ["./api"]