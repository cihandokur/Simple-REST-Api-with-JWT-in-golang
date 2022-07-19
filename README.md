# Simple REST Api with JWT Authentication, written in Go
In this repository if you will find an example of REST Api written in Go language and performing simple operations. You can find basic JWT token generation and examples using the generated token.
You can also find simple test methods.

## Technologies
    - go 1.18
    - postgres
    - docker

## Libraries used in Go
    - github.com/golang-jwt/jwt/v4 
	- github.com/gorilla/mux 
	- github.com/onsi/ginkgo/v2 
	- github.com/onsi/gomega 
	- github.com/spf13/viper
	- golang.org/x/crypto
	- gorm.io/driver/postgres
	- gorm.io/gorm

## What to do when you run it for the first time ?
Assuming you have Go and Docker installed
- for ginkgo

    > installation
    ```
    go env -w GO111MODULE=on
    go install github.com/onsi/ginkgo/v2/ginkgo@latest
    go get github.com/onsi/gomega/...
    ```
    > path update 
    ```
    <!-- go env GOPATH ->  /home/user/go -->
    export PATH=$PATH:$(go env GOPATH)/bin
    ```
    > create ginkgo suite file
    ```
    ginkgo bootstrap
    ```
    > run ginkgo test -  under the test path like /controller
    ```
    ginkgo
    - or
    go test
    ```
- env settings / configs

    You can find application settings in `/config/config.toml`.
    The following files are used for Docker settings.
    ```
    - /Dockerfile
    - /docker-compose.yml
    - /docker-entrypoint.sh
    - /.env - Please do not forget to update the POSTGRES_LOCAL_DATA_PATH variable in this file according to the operating system you are using.
        - example for mac / linux : /sql/data
        - example for windows : C:/Users/coder/sql/data"
    ```

- for docker
    > run docker compose
    ```
    docker compose up --build
    - or
    docker compose up
    ```
- for local env
    ```
    - first time
        - clone the repository to your local
        - install `go` binary or building from source https://go.dev/dl/
        - install `postgresql` binary or building from source https://www.postgresql.org/download/
        -  go mod tidy
    
    - later on
    go run cmd/main.go
    ```


### `Urls`
    
- SignUp : 

    <http://localhost:9090/signup> | POST
    - Payload
    ```json
        {
        "email": "cihan@codex11.com",
        "password": "p@zw0rd",
        "firstName": "Cihan",
        "lastName": "Dokur"
        }
    ```
    - Response
    ```json
        {
        "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNpaGFuZG9rdXJAZ21haWwuY29tIiwiZXhwIjoxNjU3ODgxODQ2fQ.yKkstBJsELTVNw9ohUt3SeC9lk0p6kt2L1XlANvKJAFpsvxE_dKpSaqWVy_U0FZo0wtDi3zlt5T3JTSt-8Jy7g" 
        }
    ```

- Login : 

    <http://localhost:9090/login> | POST
    - Payload
    ```json
        {
        "email": "cihan@codex11.com",
        "password": "p@zw0rd"
        }
    ```
    - Response
    ```json
        {
        "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNpaGFuZG9rdXJAZ21haWwuY29tIiwiZXhwIjoxNjU3ODgxODQ2fQ.yKkstBJsELTVNw9ohUt3SeC9lk0p6kt2L1XlANvKJAFpsvxE_dKpSaqWVy_U0FZo0wtDi3zlt5T3JTSt-8Jy7g" 
        }
    ```    

- Get Users : 

    <http://localhost:9090/users> | GET - `x-authentication-token key and valid token information must be included in the header.`
    - Payload
    ```
        No payload
    ```
    - Response
    ```json
        {
            "users": [
                {
                    "email": "dev@codex11.com",
                    "firstname": "Dev",
                    "lastname": "Codex"
                },
                {
                    "email": "cihan@codex11.com",
                    "firstname": "Cihan",
                    "lastname": "Dokur"
                }
            ]
        }
    ```

- Update Current User : 
    
    <http://localhost:9090/update> | PUT - `x-authentication-token key and valid token information must be included in the header.`
    - Payload
    ```json
        {
            "firstName": "Cihan",
            "lastName": "Docker"
        }
    ```
    - Response
    ```json
        - Empty if valid token information exists and transaction is successful.
        - If there is no valid token information
        {
            "Message": "invalid token.",
            "Status": "403"
        }
    ```