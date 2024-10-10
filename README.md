# Simple and well-structured API for multiple projects

This repository contains a simple API written in Go, designed to be reusable and easily integrated into multiple projects. The API is built with a focus on performance, scalability, and maintainability. 

It provides a clean and consistent interface for accessing and managing data, making it an ideal choice for a variety of applications.


# Key features

**- Good performance:** The API is designed to handle a high volume of requests with low latency.

**- Scalable:** The API can be easily scaled to accommodate increasing workloads.

**- Well-structured:** The API is well-documented and easy to use, with a consistent and intuitive interface.

**- Reusable:** The API can be easily integrated into multiple projects, reducing development time and effort.


# Use cases

**- Web applications:** The API can be used to power web applications that need to access and manage data efficiently.

**- Mobile applications:** The API can be used to develop mobile applications that need to interact with a backend server.


# To get started with the API, follow these steps:

### 1. Requirements

  - Make installed for shortcuts

  - Docker installed if you want to build and start postgres and redis containers

  - Build and start Redis container with the command ```make docker-redis```

  - Build and start postgres container with the command ```make docker-postgres```

  - Rename .env.example to ```app.env```

  - JWT .pem files with EC521 algorithm: ```keys/jwt/private.pem``` ```keys/jwt/public.pem``` 

  Others information such configurations are on ```app.env```

### 2. Clone the repository

```go
git clone https://github.com/your-username/go-api.git
```

```go
cd go-api/
```

The entry point of the project is `cmd/` folder. In this folder the is a `main.go` file.

### 3. Install dependencies

```go
make install
```

### 4. Run the API

```go
make build
```

```go
make run
```

API docs with openAPI v3.1(latest) is on 
```go
/api/v1/docs
```

Amazing API documentation(you can choose between 4 templates: Redocly, Scalar, Stoplight, Swagger). 

<ins>Scalar(default) template screenshot</ins>
![OpenAPI-Scalar](https://github.com/user-attachments/assets/0092f0e1-e2c5-4e38-a618-437097327e24)

<ins>Redocly template screenshot</ins>
![OpenAPI-Redocly](https://github.com/user-attachments/assets/1e1708aa-f355-446d-aa19-9f2ab16e08fa)

<ins>Stoplight template screenshot</ins>
![OpenAPI-Stoplight](https://github.com/user-attachments/assets/fa0595e6-46f2-48aa-a379-af19a854bc06)

<ins>Swagger template screenshot</ins>
![OpenAPI-Swagger](https://github.com/user-attachments/assets/823fbfe6-7886-450d-b58f-81b66b13f2b4)


# Features

- [x] History
  - Get history with search, filter and pagination

- [x] Role
  - CRUD operations

- [x] Role-permission
  - Create and Get all with search, filter and pagination

- [x] Auth
  - Login (📩Email, 📲Phone number, ☁️Provider['Google', 'Facebook']),
  
  - Register (📩Email, 📲Phone number),
  
  - Activate account,
    
  - Reset password.

- [x] Users
  - CRUD operations


# Contributing

We welcome contributions to this project. If you have any ideas or improvements, please feel free to open an issue or pull request.

We believe that this API can be a valuable tool for developers who need to build high-performance, scalable, and maintainable applications. We encourage you to try it out and let us know what you think.
