FROM golang:1.23.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go get -u -v github.com/golang-migrate/migrate/v4/cmd/migrate
RUN go get -u -v github.com/go-sql-driver/mysql

COPY . .

RUN go build -o main ./cmd/api

EXPOSE 8080

CMD ["./main"]