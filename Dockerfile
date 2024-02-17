FROM golang:1.21.1-bullseye as build

WORKDIR /app/bin

COPY . .

RUN go build

FROM ubuntu:22.04

COPY --from=build /app/bin/url-shortener /usr/bin/url-shortener

RUN chmod +x /usr/bin/url-shortener

ENTRYPOINT [ "url-shortener" ]