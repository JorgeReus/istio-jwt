# JWT authentication service

## Actions

### Install dependencies
1. Follow the [task install guide](https://taskfile.dev/installation/)
2. Run `task dl-deps`

### Run development mode
Run `task dev`, this will run the http service in the port `8080`

## Build docker image
Run `task docker-build`

## Local testing
1. Go to the [swagger-docs](http://localhost:8080/swagger/index.html)
2. Play around with the endpoints
3. Explore the tokens (id_tokens & refresh_tokens) in the [jwt debugger](https://jwt.io/)
