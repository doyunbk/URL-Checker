FROM golang:1.13.1 AS builder
WORKDIR $GOPATH/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/url_checker $GOPATH/main.go

# Add built binay to docker image using scratch
FROM scratch
COPY --from=builder /go/bin/url_checker /go/bin/url_checker
EXPOSE 8000
ENTRYPOINT ["/go/bin/url_checker"]
