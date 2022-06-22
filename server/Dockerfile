FROM golang:1.17

WORKDIR .

COPY go.mod go.sum server/
RUN cd server && go mod download && go mod verify

COPY . server/

RUN cd server && go mod tidy && go build main.go

CMD server/main -dsn="root:123456@tcp(172.16.1.3:3306)/dis_db?parseTime=true"
