# Notion API Server

A RESTful API server built with Go and Gin framework following Clean Architecture principles.

## Project Structure

```
notion/
├── src/
│   ├── cmd/                    # Command line interface
│   ├── config/                 # Configuration management
│   ├── middleware/             # HTTP middleware
│   ├── internal/               # Internal application code
│   │   ├── domain/            # Domain entities and interfaces
│   │   ├── repository/        # Data access layer
│   │   ├── usecase/           # Business logic layer
│   │   ├── delivery/          # HTTP handlers
│   │   └── infrastructure/    # External dependencies (database, etc.)
│   └── migrations/            # Database migrations
├── main.go                    # Application entry point
├── go.mod                     # Go modules
├── env_example               # Environment variables example
└── README.md                 # This file
```

## Features

- Clean Architecture implementation
- RESTful API with Gin framework
- MySQL database with GORM
- User authentication and registration
- Environment-based configuration
- Middleware for CORS and error handling
- Structured logging with Zap

## Prerequisites

- Go 1.24 or higher
- MySQL 8.0 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd notion
```

2. Install dependencies:
```bash
go mod tidy
```

3. Copy environment file:
```bash
cp env_example .env
```

4. Configure your environment variables in `.env` file:
```env
# Server Configuration
PORT=8080
ENV=development
LOG_LEVEL=info

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=your_password
DB_NAME=notion
```

5. Create MySQL database:
```sql
CREATE DATABASE notion;
```

## Running the Application

### Development Mode
```bash
go run main.go server
```

### Production Mode
```bash
go build -o notion main.go
./notion server
```

## API Endpoints

### Health Check
- `GET /health` - Health check endpoint

### User Management
- `POST /api/v1/users/register` - Register new user
- `POST /api/v1/users/login` - User login
- `GET /api/v1/users/:id` - Get user by ID
- `GET /api/v1/users/` - List users (with pagination)

### Example Requests

#### Register User
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "name": "John Doe"
  }'
```

#### Login User
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

#### Get User by ID
```bash
curl http://localhost:8080/api/v1/users/1
```

#### List Users
```bash
curl "http://localhost:8080/api/v1/users/?limit=10&offset=0"
```

## Clean Architecture

This project follows Clean Architecture principles:

1. **Domain Layer** (`src/internal/domain/`): Contains business entities and interfaces
2. **Repository Layer** (`src/internal/repository/`): Data access implementation
3. **Use Case Layer** (`src/internal/usecase/`): Business logic implementation
4. **Delivery Layer** (`src/internal/delivery/`): HTTP handlers and API endpoints
5. **Infrastructure Layer** (`src/internal/infrastructure/`): External dependencies like database

## Development

### Adding New Features

1. Define domain entities in `src/internal/domain/`
2. Implement repository in `src/internal/repository/`
3. Add business logic in `src/internal/usecase/`
4. Create HTTP handlers in `src/internal/delivery/http/`
5. Add routes in `src/cmd/server.go`

### Database Migrations

Database migrations are handled automatically by GORM's AutoMigrate feature. When you add new domain entities, they will be automatically migrated.

## Testing

```bash
go test ./...
```

## Docker Support

Build and run with Docker:

```bash
docker build -t notion-api .
docker run -p 8080:8080 notion-api
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.