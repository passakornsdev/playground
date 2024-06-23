FROM golang

WORKDIR /home/go-rest/

COPY . .

RUN apt-get update && \
    apt-get install -y nmon && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o go-server

EXPOSE 8080

CMD ["./go-server"]
