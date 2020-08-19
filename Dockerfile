FROM golang:1.14-alpine3.11 AS build-env

RUN apk add --no-cache git make

RUN mkdir /cloudaudition/

WORKDIR /cloudaudition/
ADD . /cloudaudition/

RUN make build

FROM alpine:3.8

WORKDIR /

COPY --from=build-env /cloudaudition/build/bin/cloudaudition /

EXPOSE 8080

CMD ["/cloudaudition"]