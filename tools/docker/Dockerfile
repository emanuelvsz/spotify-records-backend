FROM golang:1.19

WORKDIR /app

RUN mkdir tools
COPY src src
COPY go.mod go.mod

RUN apt-get update
RUN apt-get install dos2unix
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go build -o /server
RUN main.go 

EXPOSE 8080

CMD ["/server"]