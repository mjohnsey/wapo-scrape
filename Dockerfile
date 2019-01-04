FROM golang:1.11.1-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on
RUN apk add --no-cache git

WORKDIR /go/src/github.com/mjohnsey/wapo-scrape
COPY . /go/src/github.com/mjohnsey/wapo-scrape
RUN go install -ldflags '-s -w'

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/wapo-scrape /wapo-scrape
ENTRYPOINT ["./wapo-scrape"]
