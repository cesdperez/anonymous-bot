FROM golang:1.11.3
ENV GOBIN /go/bin
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get .
RUN go build -o main .
CMD ["/app/main"]
