FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./main.go ./
COPY ./app/ ./app
COPY ./*.env ./

RUN go build -o /goapp ./

EXPOSE 8080

CMD ["/goapp"]