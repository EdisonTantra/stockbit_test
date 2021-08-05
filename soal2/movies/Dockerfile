# FROM alpine
# ADD movies /movies

FROM golang:1.15.14-alpine3.13 as builder

RUN apk update && apk add git && apk add ca-certificates

ENV GO111MODULE=on

RUN mkdir /src
WORKDIR /src

 # <- COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o movie .

# Second step to build minimal image
FROM alpine

RUN apk add -U --no-cache tzdata bash
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && echo "Asia/Jakarta" > /etc/timezone
RUN date

RUN export TZ="Asia/Jakarta"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /src/movie /movie
