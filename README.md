# Simple and well-structured API for multiple projects

This repository contains a simple API written in Go, designed to be reusable and easily integrated into multiple projects. The API is built with a focus on performance, scalability, and maintainability. 

It provides a clean and consistent interface for accessing and managing data, making it an ideal choice for a variety of applications.


# Key features

**- Good performance:** The API is designed to handle a high volume of requests with low latency.

**- Scalable:** The API can be easily scaled to accommodate increasing workloads.

**- Well-structured:** The API is well-documented and easy to use, with a consistent and intuitive interface.

**- Reusable:** The API can be easily integrated into multiple projects, reducing development time and effort.

# To get started with the API, follow these steps:

### 1. Requirements

  - Make installed for shortcuts

  - Docker installed if you want to build and start postgres or redis containers

  - Build and start Redis container with the command ```make docker-redis```

  - Build and start postgres container with the command ```make docker-postgres```

  - Rename .env.example to ```app.env```

  - JWT .pem files with ES212(ECDSA SHA-512) algorithm: ```keys/jwt/private.pem``` ```keys/jwt/public.pem```
    You ca use this website to generate JWT keys for your tests [JWT online generator](https://jwt-keys.21no.de/) 

  - Password is hashed using Argon2id algorithm. If you want to customize salinity, you can edit the .env.example file


  Others information such configurations are on ```app.env```

### 2. Clone the repository

```
git clone https://github.com/your-username/go-api.git
```

```
cd go-api/
```

The entry point of the project is `cmd/` folder. In this folder the is a `main.go` file.

### 3. Install dependencies

```
make install
```

### 4. Run the API

```
make build
```

```
make run
```

API docs with openAPI v3.1(latest) is on 
```
/api/v1/docs
```

If you want to scan vulnerabilities(security issues)
```
make scan
```

# Features

- [x] History

- [x] Roles

- [x] Permissions

- [x] Authentication

- [x] Users

# TODO

- Add 2FA and passKey, and add Google reCAPTCHA to all endpoints starting with /auth

- Add testing

- Add CI pipeline with GitHub actions(build image)


Amazing API documentation. You can choose between 4 templates: Scalar(Default), Redocly, Stoplight, Swagger. 

<ins>Scalar(default) template screenshot</ins>
![OpenAPI-Scalar](https://github.com/user-attachments/assets/0092f0e1-e2c5-4e38-a618-437097327e24)

<ins>Redocly template screenshot</ins>
![OpenAPI-Redocly](https://github.com/user-attachments/assets/1e1708aa-f355-446d-aa19-9f2ab16e08fa)

<ins>Stoplight template screenshot</ins>
![OpenAPI-Stoplight](https://github.com/user-attachments/assets/fa0595e6-46f2-48aa-a379-af19a854bc06)

<ins>Swagger template screenshot</ins>
![OpenAPI-Swagger](https://github.com/user-attachments/assets/823fbfe6-7886-450d-b58f-81b66b13f2b4)


# Contributing

