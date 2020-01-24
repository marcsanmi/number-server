#FROM golang:alpine as builder
#
## Environments defaults
#ENV PORT 4000
#ENV PATH ./cmd/server/main.go
#ENV GO111MODULE=on
#
##RUN apk update && apk add bash ca-certificates git
#RUN apk add --no-cache bash
#
## Set working directory
#WORKDIR /newrelic
#
## Copy the current code into our workdir
#COPY . .
#
#RUN go mod download
#
## Build go binary with some flags in order to run it in Alpine
#RUN CGO_ENABLED=0 GOOS=linux go build -o newrelic -a -installsuffix cgo $PATH
#
#FROM alpine:latest
#
#RUN apk --no-cache add ca-certificates
#
## Pull the binary from the builder container
#COPY --from=builder /newrelic .
#
## Run the container
#CMD ["./newrelic"]

FROM golang:alpine as builder

RUN apk --no-cache add git

#ENV PATH=./cmd/server/main.go

ARG path="./cmd/client/main.go"

WORKDIR /newrelic

COPY . .

RUN go mod download

#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
RUN CGO_ENABLED=0 GOOS=linux go build -o newrelic -a -installsuffix cgo $path

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /newrelic .

ENTRYPOINT ["./newrelic"]