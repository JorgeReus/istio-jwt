FROM golang:1.16-alpine as build
COPY src /src
WORKDIR /src
RUN go mod download && go build -o app

FROM alpine:latest
WORKDIR /root/
COPY --from=build /src/ .
EXPOSE 8080
CMD ["/root/app"]
