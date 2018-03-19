FROM golang:1.10

WORKDIR /go/src/bmp085
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# FROM resin/raspberry-pi2-alpine:3.4
# WORKDIR /app
# COPY --from=0 /go/bin/bmp085    .

EXPOSE 8080

CMD ["bmp085"]
# CMD ["/app/bmp085"]
