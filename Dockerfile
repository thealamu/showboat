FROM golang:alpine

WORKDIR /app
COPY . .

RUN go get -v
RUN go build -o showboat

CMD ["./showboat"]

