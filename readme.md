# assignment-totality-corp

This repository hosts a gRPC server project for Totality Corp.

## Overview

This project exemplifies a gRPC server setup tailored for a Totality Corp assignment. It showcases the process of creating and operating a gRPC server using Go, with guidance for deployment in both Docker and non-Docker environments.

## Features

- Developed a gRPC server using Go
- Configured to run on port 8080 by default
- Includes multiple gRPC services and endpoints as specified in the assignment

## Prerequisites

Ensure the following tools are installed before running the server:

- [Go](https://golang.org/dl/) (version 1.20 or higher)
- [Docker](https://www.docker.com/get-started) (for Docker-based deployment)

## Usage

### Using Docker

- **To build a Docker image:**

    ```bash
    make build
    ```

- **To start the server:**

    ```bash
    make run
    ```

    By default, the server will listen on port 8080.

### Without Docker

- **To start the server without Docker:**

    ```bash
    go run cmd/server/userServer.go
    ```

    By default, the server will listen on port 8080.

## Testing

- **To execute tests:**

    ```bash
    make test
    ```

- **To execute tests with coverage reporting:**

    ```bash
    make coverage
    ```

- **To test using a client, first start the server and then run:**

    ```bash
    go run client/client.go
    ```