# Food Store Calculator

Food Store Price Calculator implemented in Go following Clean Architecture principles (cmd -> handler -> service -> model).

## Project Structure
- `cmd/main.go`: Application entry point
- `internal/handler`: Handles checkout requests and responses
- `internal/service`: Contains business logic for bundle & member card discounts
- `internal/model`: Defines menu items, prices, and data transfer objects

## How to Run

Run the main application:
```bash
go run cmd/main.go
```

Run unit tests:

```bash
go test -v ./internal/service/...
```