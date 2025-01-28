FROM golang:1.20

WORKDIR /app

COPY *.go ./

RUN go mod init log-output && go mod tidy

RUN go build -o ./log-output

CMD ["./log-output"]

