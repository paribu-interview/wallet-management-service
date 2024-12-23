# Wallet Management Service

The Wallet Management Service (WMS) is a microservice designed for managing wallets. It provides endpoints for creating,
retrieving, and deleting wallets.

## Features

- Create a new wallet with a unique address and network.
- Retrieve wallet details by ID.
- Delete wallets by ID.

## Requirements

- Go 1.19 or later
- PostgreSQL
- Docker and Docker Compose (optional for containerized deployment)

## Setup and Installation

### Clone the Repository

```bash
git clone https://github.com/yourusername/wallet-management-service.git
cd wallet-management-service
```

## Run Locally

### Set Up the Database

1. Create a new PostgreSQL database.
2. Update the database connection details in the `local.env` file.

   ```env 
   PG_USER=your_db_user
   PG_PASSWORD=your_db_password
   PG_NAME=your_db_name
   PG_HOST=your_db_host
   PG_PORT=your_db_port
   ```

3. Set APP_ENV to `local`.

    ```env
    export APP_ENV=local
    ```
4. Install the dependencies.
   ```bash
   go mod tidy
   ```
5. Run the server.
   ```bash
   go run cmd/main.go
   ```

## Run with Docker

1. Build and run the Docker container.
   ```bash
   docker-compose up --build
   ```

2. The application will start on `http://localhost:8080/api`.

## Migration

Migration is automatically handled by the GORM library. The database schema is created when the application starts.

## API Endpoints

The following endpoints are available:

- `POST /api/wallets`: Create a new wallet.
- `GET /api/wallets/{id}`: Retrieve wallet details by ID.
- `DELETE /api/wallets/{id}`: Delete a wallet by ID.

### Create a new wallet:

- Request:

   ```http
   POST /api/wallets
   Content-Type: application/json
   ```
- Request Body:
  ```json
  {
    "address": "0x1234567890abcdef1234567890abcdef12345678",
    "network": "ethereum"
  }
  ```

- Response Body:

   ```json
   {
    "id": 1,
    "address": "0x1234567890abcdef1234567890abcdef12345678",
    "network": "ethereum"
   }
   ```
- Response
    - 201 Created: Wallet created successfully.
    - 400 Bad Request: Invalid input.
    - 409 Conflict: Wallet already exists.

### Retrieve wallet details by ID:

- Request:

   ```http
   GET /api/wallets/1
   ```
- Response Body:

   ```json
   {
    "id": 1,
    "network": "ethereum",
    "address": "0x1234567890abcdef1234567890abcdef12345678"
   }
   ```
- Response
    - 200 OK: Wallet details retrieved successfully.
    - 404 Not Found: Wallet not found.
    - 400 Bad Request: Invalid input.
    - 500 Internal Server Error: Server error.

### Delete a wallet by ID:

- Request:

  ```http
  DELETE /api/wallets/1
  ```
- Response Body:

    ```json
    {
    "message": "Wallet deleted successfully."
    }
    ```
- Response
    - 204 No Content: Wallet deleted successfully.
    - 404 Not Found: Wallet not found.
    - 400 Bad Request: Invalid input.
    - 500 Internal Server Error: Server error.

## Testing

Run the tests using the following command:

```bash
go test ./...
```
