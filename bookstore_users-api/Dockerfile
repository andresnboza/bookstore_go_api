FROM golang:1.16-alpine AS build
WORKDIR /src
COPY . /src
RUN go build -o bookstore_users_bin main.go

FROM alpine
COPY --from=build /src/bookstore_users_bin /
ENTRYPOINT ["/bookstore_users_bin"]