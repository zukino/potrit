# Potrit Social Media REST API

A REST API for a social media application built with Go and PostgreSQL following Clean Architecture principles.

## Development Setup

### Prerequisites
- Go 1.21+ installed
- PostgreSQL 14+ installed and running
- Git for version control

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd potrit
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up the database:
```bash
# Create database
createdb potrit
```

4. Run the application:
```bash
go run cmd/server/main.go
```

## Dependencies

- **PostgreSQL Driver**: `github.com/lib/pq` - For database connectivity
- **JWT Library**: `github.com/golang-jwt/jwt/v5` - For token-based authentication

## Architecture

This project follows Clean Architecture principles with clear separation of concerns:

- **Domain Layer**: Core business entities and repository interfaces
- **Application Layer**: Use cases and business logic
- **Infrastructure Layer**: External concerns (database, authentication)
- **Presentation Layer**: HTTP handlers and routing

## API Endpoints

The API will provide endpoints for:
- User registration and authentication
- Post creation and management
- Connection management
- Content interactions (likes, comments)

## License

MIT License