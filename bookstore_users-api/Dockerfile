# FROM golang:1.16-alpine AS build
# WORKDIR /src
# COPY . /src
# RUN go build -o bookstore_users_bin main.go

# FROM alpine
# COPY --from=build /src/bookstore_users_bin /
# EXPOSE 8080
# ENTRYPOINT ["/bookstore_users_bin"]


# Having hot reload

FROM golang:latest

RUN go get -v github.com/canthefason/go-watcher/cmd/watcher

WORKDIR /go/src/github.com/andresnboza/bookstore_users-api

COPY . .

ENTRYPOINT /go/bin/watcher