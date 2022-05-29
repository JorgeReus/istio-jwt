FROM golang:alpine as base
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o authz-service ./cmd/web/main.go

FROM alpine:latest
WORKDIR /app
EXPOSE 8080
ENV USER=authz-service
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --no-create-home \
    "$USER"
USER "$USER"
COPY --from=base /app/authz-service .
CMD ["./authz-service"]
