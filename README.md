# Simple and well-structured API for multiple projects

This repository contains a simple API written in Go, designed to be reusable and easily integrated into multiple
projects. The API is built with a focus on performance, scalability, and maintainability.

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

# To get started with the API, follow these steps:

### 1. Requirements

- Make installed for shortcuts

- Docker installed if you want to build and start postgres or redis containers

- Build and start Redis container with the command ```make docker-redis```

- Build and start postgres container with the command ```make docker-postgres```

- Rename .env.example to ```app.env```

- JWT .pem files with ES512(ECDSA SHA-512) algorithm: ```./assets/private/keys/jwt/private.pem``` ```./assets/private/keys/jwt/public.pem```
  You ca use this website to generate JWT keys for your tests [JWT online generator](https://jwt-keys.21no.de/)

- Password is hashed using Argon2id algorithm. If you want to customize salinity, you can edit the .env.example file

Others information such configurations are on ```app.env```

### 2. Clone the repository

```
git clone https://github.com/4kpros/go-api.git
```

```
cd go-api/
```

The entry point of the project is `cmd/` folder. In this folder the is the `main.go` file.

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

You can choose between 4 templates: Scalar(Default), Redocly, Stoplight, Swagger.

<ins>Scalar(default) template screenshot</ins>
![Scalar](https://github.com/user-attachments/assets/b0a5304a-7b77-496c-a5c8-d6f81581aba5)

<ins>Redocly template screenshot</ins>
![Redocly](https://github.com/user-attachments/assets/abb7a0b4-e481-4a8e-8fba-08498283ad21)

<ins>Stoplight template screenshot</ins>
![Stoplight](https://github.com/user-attachments/assets/2b18f7f3-2577-4617-b64a-b5981de3dfc3)

<ins>Swagger template screenshot</ins>
![Swagger](https://github.com/user-attachments/assets/9f2a7fab-4472-42f7-bad6-33d46b44f374)

