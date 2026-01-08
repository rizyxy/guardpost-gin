# Guardpost Gin

Guardpost Gin is a powerful and lightweight API Gateway and Reverse Proxy built with Go and the Gin Web Framework. It features dynamic routing configuration, a built-in authentication system, and live-reloading support for development.

## 🚀 Features

- **Dynamic Reverse Proxy**: Configure routes and downstream services via `routes.yaml`.
- **Authentication System**: Built-in support for User Registration, Login, and JWT-based authentication.
- **Refresh Token Support**: Secure token refresh mechanism to maintain user sessions.
- **Live Reloading**: Seamless development experience using [Air](https://github.com/air-verse/air).
- **API Documentation**: Pre-configured [Bruno](https://usebruno.com/) collections for easy testing.

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
- **ORM**: [GORM](https://gorm.io/)
- **Security**: JWT (JSON Web Tokens), bcrypt for password hashing
- **Configuration**: YAML, Dotenv

## 📋 Prerequisites

- [Go](https://go.dev/dl/) 1.25 or higher
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) (optional, for containerized deployment)
- [Air](https://github.com/air-verse/air) (optional, for live reloading)

## ⚙️ Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/guardpost-gin.git
   cd guardpost-gin
   ```

2. **Setup environment variables:**

   ```bash
   cp .env.example .env
   ```

   Edit `.env` and set your `JWT_SECRET` and `DB_URL` (usually `test.db`).

3. **Configure routes:**
   ```bash
   cp routes-example.yaml routes.yaml
   ```
   Modify `routes.yaml` to define your proxy routes.

## 🚀 Running the Application

### Using Makefile (Recommended)

The project includes a `Makefile` for common tasks:

```bash
make build          # Build the binary
make run            # Build and run the gateway
make docker-up      # Start the gateway using Docker Compose
make docker-down    # Stop the Docker containers
make clean          # Remove binary and data directory
```

### Development (with Air)

```bash
air
```

### Using Docker Compose

```bash
docker-compose up --build -d
```

### Manual Run

```bash
go mod tidy
go run main.go
```

## 📖 Configuration

### Dynamic Routes (`routes.yaml`)

Define your proxy routes in `routes.yaml`:

```yaml
routes:
  - path: /example
    method: GET
    downstream: https://api.example.com
```

### Environment Variables (`.env`)

- `DB_URL`: Path to your database file.
- `JWT_SECRET`: Secret key for signing tokens.

## 📂 Project Structure

- `config/`: Configuration loaders (DB, Env, Routes).
- `internal/`: Core logic and internal API routes.
- `controllers/`: Auth and request handling logic.
- `models/`: GORM database models.
- `repository/`: Data access layer.
- `routes/`: Routing logic for the auth system.
- `bruno/`: API collection for Bruno.
- `Dockerfile`: Container image definition.
- `docker-compose.yaml`: Multi-container orchestration.
- `Makefile`: Shortcut commands for development.
- `main.go`: Application entry point.

## 🧪 Testing with Bruno

Import the collection found in the `bruno/` directory into [Bruno](https://usebruno.com/) to start testing the authentication and proxy endpoints.

## 📄 License

This project is licensed under the MIT License.
