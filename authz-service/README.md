# JWT authorization service

This is an example GRPC & HTTP service that will decode a JWT in raw form(bas64) and will parse the claim, groups. 
If you have a string `"admin"` in this claim your request will be allowed, permission denied otherwhise.

## Actions

### Install dependencies
1. Follow the [task install guide](https://taskfile.dev/installation/)
2. Run `task dl-deps`

### Run development mode
Run `task dev`, this will run the http service in the port `8080` & the GRPC service in the port `9090`

## Build docker image
Run `task docker-build`

## Local testing
1. Generate a jwt token using the [auth service](../auth-service) or something like [this](http://jwtbuilder.jamiekurtz.com/)
2. Send an http request with a JWT authorization header (e.g. `http localhost:8080 "Authorization: Bearer $jwt"`)
