Food Store Price Calculator implemented in Go following Clean Architecture principles (cmd -> handler -> service -> model).

## Project Structure
- `cmd/main.go`: Application entry point and HTTP Server setup
- `internal/handler`: Handles checkout HTTP requests and responses
- `internal/service`: Contains business logic for bundle & member card discounts
- `internal/model`: Defines menu items, prices, and data transfer objects

## How to Run

### 1. Run the Main Application
Running `main.go` will execute default test cases from the assignment via Console output and automatically start the HTTP API Server at `http://localhost:8080`:

```bash
go run cmd/main.go
```

### 2. Run Unit Tests

```bash
go test -v ./internal/service/...
```

You can test custom orders via HTTP POST requests:
Endpoint: POST http://localhost:8080/checkout

cURL Request:

```bash
curl -X POST http://localhost:8080/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "items": {
      "Red": 1,
      "Green": 1
    },
    "has_member_card": true
  }'
```

Response:

```json
{
  "total": 81
}
```