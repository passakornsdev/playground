FROM golang

WORKDIR /home/go-rest/

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o go-server

EXPOSE 8080

CMD ["./go-server"]
