FROM golang
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY words /words

RUN CGO_ENABLED=0 GOOS=linux go build -o /jumble

EXPOSE 8080

CMD ["/jumble"]