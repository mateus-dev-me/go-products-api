# Documentação do Projeto

![Go Versions](https://img.shields.io/badge/go-v1.23.0-blue) ![License](https://img.shields.io/badge/license-MIT-green)

> A Simple API built using Clean Architecture.

## **Table of Contents**
1. [Project Overview](#project-overview)
2. [Project Structure](#project-structure)
3. [Installation](#installation)
4. [Usage](#usage)

## **Project Overview**

This project is a web application designed to provide an API for managing products in a system. The project is structured based on Clean Architecture, promoting a modular and maintainable organization that facilitates upkeep and expansion.

## **Project Structure**

The project structure is organized as follows:

```
.
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── domain/
│   ├── infrastructure/
│   └── use_cases/
│
├── interfaces/
│   └── api/
│       └── controllers/
│
├── tests/
│   ├── use_cases/
│   ├── infrastructure/
│   └── controllers/
│
├── migrations/
│
├── docs/
│
├── go.mod
└── go.sum
```


## **Installation**

### **Prerequisites:**

- **Go** (version 1.23.0 or higher)
- **Docker** (for container setup)

1. Clone the repository:

    ```bash
    git clone https://github.com/mateus-dev-me/ama-api.git
    cd ama-api
    ```

2. Configure environment variables:

    Create a `.env` file based on the `.env.example` file and configure your database credentials and parameters.

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Run database migrations:

    ```bash
    go generate
    ```

## **Usage**

### **Start the Server:**

```bash
go run ./cmd/main.go
```

The API will be available at http://localhost:8080/api/v1

