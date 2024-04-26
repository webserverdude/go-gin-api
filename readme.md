# Gin API for Managing Gin Spirits

This is a simple API written in Go using the Gin web framework. It allows you to manage a collection of gin spirits, including retrieving all gins, adding new gins, and fetching a specific gin by its ID.

## Getting Started

To get started with using this API, follow the instructions below:

### Prerequisites

- Go installed on your system. You can download and install it from [golang.org](https://golang.org/doc/install).
- Git installed on your system. You can download and install it from [git-scm.com](https://git-scm.com/downloads).

### Installation

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/webserverdude/go-gin-api.git
    ```

2. Change into the directory:

    ```bash
    cd gin-api
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

### Creating SSL Keys

To enable HTTPS for your server, you'll need to create SSL private and public keys. Here's how to do it:

1. Open a terminal window.

2. Run the following command to generate a private key file (`server.key`):

    ```bash
    openssl genrsa -out server.key 2048
    ```

3. Run the following command to generate a certificate signing request (CSR):

    ```bash
    openssl req -new -key server.key -out server.csr
    ```

4. Now, create a self-signed certificate using the CSR and private key:

    ```bash
    openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
    ```

    This command will create a self-signed certificate (`server.crt`) valid for 365 days.

### Usage

1. Run the Go server:

    ```bash
    go run main.go
    ```

2. The API server will start running on `https://localhost:8000`. You can now make requests to the following endpoints:

    - `GET /gins`: Retrieve all gins.
    - `GET /gins/{id}`: Retrieve a specific gin by its ID.
    - `POST /gins`: Add a new gin.

3. You can use tools like cURL or Postman to interact with the API endpoints.

## API Documentation

The API documentation is provided using an OpenAPI specification file. You download the file by navigating to `https://localhost:8000/openapi.json` after starting the server.

## License

This project is licensed under the [MIT License](LICENSE).
