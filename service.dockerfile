FROM golang:alpine AS base_build
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

COPY . /go/src/github.com/mkuchenbecker/storage/
WORKDIR /go/src/github.com/mkuchenbecker/storage/


RUN go get ./...
RUN apk add ca-certificates


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags='-w -s -extldflags "-static"' -o /go/bin/service ./service

FROM golang:alpine as slate
# # Copy our static executable.
COPY --from=base_build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base_build /go/bin/service /go/bin/service

ENTRYPOINT ["/go/bin/service"]
EXPOSE 9100:9109
