FROM alpine:latest
WORKDIR /app
EXPOSE 8080
ENV USER=auth-service
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --no-create-home \
    "$USER"
USER "$USER"
COPY ./auth-service .
CMD ["./auth-service"]
