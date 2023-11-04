FROM golang
WORKDIR /app

COPY go.mod words *.html *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /jumble

EXPOSE 8080

CMD ["/jumble"]