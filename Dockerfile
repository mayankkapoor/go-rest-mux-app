# build stage
FROM golang:1.13-alpine AS build-env
RUN apk add --no-cache git mercurial \
   && go get github.com/go-sql-driver/mysql \
   && go get github.com/gorilla/mux \
   && apk del git mercurial

WORKDIR /usr/src/app

COPY . .
RUN go build -o goapp

# final stage: Create image with only build artifacts
FROM alpine
WORKDIR /usr/local/bin
COPY --from=build-env /usr/src/app/goapp /usr/local/bin/

EXPOSE 8000
ENTRYPOINT ./goapp
