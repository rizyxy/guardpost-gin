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

- [Go](https://go.dev/dl/) 1.21 or higher
- [Air](https://github.com/air-verse/air) (optional, for live reloading)

## ⚙️ Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/guardpost-gin.git
   cd guardpost-gin
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Setup environment variables:**

   ```bash
   cp .env.example .env
   ```

   Edit `.env` and set your `JWT_SECRET` and `DB_URL` (usually `test.db`).

4. **Configure routes:**
   ```bash
   cp routes-example.yaml routes.yaml
   ```
   Modify `routes.yaml` to define your proxy routes.

## 🚀 Running the Application

### Development (with Air)

```bash
air
```

### Production

```bash
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

- `DB_URL`: Path to your SQLite database.
- `JWT_SECRET`: Secret key for signing tokens.

## 📂 Project Structure

- `config/`: Configuration loaders (DB, Env, Routes).
- `internal/`: Core logic and internal API routes.
- `controllers/`: Auth and request handling logic.
- `models/`: GORM database models.
- `repository/`: Data access layer.
- `routes/`: Routing logic for the auth system.
- `bruno/`: API collection for Bruno.
- `main.go`: Application entry point.

## 🧪 Testing with Bruno

Import the collection found in the `bruno/` directory into [Bruno](https://usebruno.com/) to start testing the authentication and proxy endpoints.

## 📄 License

This project is licensed under the MIT License.
