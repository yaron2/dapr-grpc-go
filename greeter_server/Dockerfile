FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go mod tidy
RUN go build -o /server

CMD [ "/server" ]
