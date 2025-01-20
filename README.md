# Book microservice architecture

## Start microservices

### Run api-gateway

```
cd ./api-gateway
go run main.go
```

### Run auth-service

```
cd ./auth-service
go run main.go
```

### Run resources-service

```
cd ./resource-service
go run main.go
```

### Run api-gateway

## Available endpoints

## Auth

### Registration POST http://localhost:8080/api/v1/sign_in

Request body:
```
{
    "username":"kyrkela",
    "password":"your_password"
}
```
Response:
```
"access_token": "your_access_token",
    "refresh_token": "your_refresh_token"
```

## Resources

### GetBooks GET http://localhost:8080/api/v1/users/sign_in
Response:
```
[
    {
        "author": "J.K.Rowling",
        "title": "Harry Potter 1"
    },
    {
        "author": "J.K.Rowling",
        "title": "Harry Potter 2"
    },
    {
        "author": "J.R.R. Tolkien",
        "title": "Lord of the Rings"
    }
]
```

### GetUsers GET http://localhost:8080/api/v1/users

#### NOTE - user must have Authorization header with valid access_token

Response:
```
[
    {
        "username": "kyrkela",
        "email": "volodymyrshabat@gmail.com",
        "roles": [
            "user",
            "vip"
        ]
    },
    {
        "username": "avbyte",
        "email": "nazardubenko2@gmail.com",
        "roles": [
            "user"
        ]
    },
    {
        "username": "yulianam",
        "email": "yulianamal@gmail.com",
        "roles": [
            "user",
            "admin"
        ]
    }
]
```