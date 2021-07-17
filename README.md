# An example of clean microservice

## Local installation

Create vendor directory

```sh
go mod vendor
```

## Commands

### HTTP server

```sh
go run ./cmd/httpserver --help
```

## Testing

Running unit tests

```sh
go test ./internal/...
```

Running functional tests

```sh
go test ./functests/...
```

Running all tests on the project

```sh
go test ./...
```

## Swagger

Load the swagger UI docker image

```sh
docker pull swaggerapi/swagger-editor
```

Running HTTP server with rendered Swagger UI 

_`bash` command, for `fish` terminal remove `$` sign_

```bash
docker run -d --rm -p 80:8080 -v $(pwd)/swagger:/tmp -e SWAGGER_FILE=/tmp/swagger.yml swaggerapi/swagger-editor
```

Go to http://127.0.0.1
