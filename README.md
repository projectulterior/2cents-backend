# 2CENTS BACKEND
This repository contains the GraphQL backend server for 2cents written in Golang.

## Running
```sh
make up
```

## Authentication (JWT)
- Create Token
    ```sh
    curl -X POST http://localhost:8080/auth/create_token \
        --header 'Content-Type: application/json' \
        --data '{ "username": "username", "password": "password1" }'
    ```

# Resources

- [gqlgen](https://gqlgen.com/getting-started/)