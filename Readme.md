# How to running 101

## Setup Go Environment

1. create go module path

```
go mod init cobain/cobain-restapi
go mod tidy
```

2.installing dependencies

```
go get .
```

3. running the server

```
go run .
```

## Test the API in localhost:8080

1. Get request

```
curl http://localhost:8080/albums
```

2. Add json file from main.json (POST request)

```
curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"

```
