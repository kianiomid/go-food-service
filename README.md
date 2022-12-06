<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Food Service

## About the service

The service is used to create about 10 APIs that included User and Food Models.

## Clone the project

```
$ git clone https://github.com/kianiomid/go-food-service
$ cd go-food-service
```

## Getting started

Below we describe the conventions or tools specific to golang project.

### Layout

```tree
├── cmd
│   └── main.go
├── configurations
│   └── App.yaml
├── data
│   └── migrations
│       └── food-app.sql
├── internal
│   ├── admin
│   │   └── dto
│   │   └── entity
│   │   └── handler
│   │   └── repository
│   │   │    └── repositoryInterfaces
│   │   └── rule
│   │   └── service
│   │   │   └── serviceInterfaces
│   │   └── transformer
│   └── presentation
│       └── middleware
│       └── route
└── pkg
│   ├── config
│   ├── database
│   ├── jwttoken
│   ├── response
│   └── security
├── Dockerfile
├── go.mod
├── go.sum
└── README.me
```

## Run the project
### 1. Build docker image
```
$ docker build -t food-service . --target production
```
### 2. Run food-service image as food-service-container
```
$ docker run -d -p <your-custom-port>:8090 --name food-service-container food-service
```