# by_te — URL Shortener

A URL shortener service built in Go, with MySQL for persistence, Redis for caching, and Nginx for load balancing across two app instances.

---

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.22+](https://go.dev/dl/) _(only needed for local development)_

---

## Installation & Setup

### Option 1: Docker (recommended)

1. **Clone the repository**

   ```bash
   git clone <repo-url>
   cd by_te
   ```

2. **Build and start all services**

   ```bash
   docker compose -f deployments/docker/docker-compose.yml up --build
   ```

   This starts:
   | Service | URL |
   |---------|-----|
   | Nginx (load balancer) | http://localhost:8080 |
   | App instance 1 | http://localhost:8081 |
   | App instance 2 | http://localhost:8082 |
   | Adminer (DB UI) | http://localhost:8083 |
   | MySQL | localhost:3307 |
   | Redis | localhost:6379 |

3. **Verify the service is running**

   ```bash
   curl http://localhost:8080/health
   # OK
   ```

4. **Stop all services**

   ```bash
   docker compose -f deployments/docker/docker-compose.yml down
   ```

   To also remove the MySQL data volume:

   ```bash
   docker compose -f deployments/docker/docker-compose.yml down -v
   ```

---

### Option 2: Local Development

1. **Clone the repository**

   ```bash
   git clone <repo-url>
   cd by_te
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Start MySQL and Redis**

   ```bash
   docker compose -f deployments/docker/docker-compose.yml up -d redis mysql
   ```

4. **Configure environment**

   The app loads `.env.<APP_ENV>` (defaults to `local`), so `.env.local` is used automatically. Update the values for local development:

   ```env
   DB_HOST=localhost
   DB_PORT=3307   # host-mapped port from Docker
   REDIS_ADDR=localhost:6379
   ```

5. **Run the application**

   ```bash
   go run ./cmd/api
   ```

6. **Verify**

   ```bash
   curl http://localhost:8080/health
   # OK
   ```

---

## Project Structure

```
by_te/
├── cmd/api/                  # Entry point
├── internal/
│   ├── config/               # Environment config loader
│   ├── domain/url/           # Domain models, interfaces
│   ├── application/url/      # Use cases
│   ├── infrastructure/
│   │   ├── cache/redis/      # Redis client
│   │   ├── database/mysql/   # MySQL client + repository
│   │   ├── idgenerator/      # Snowflake ID generator
│   │   └── logger/           # Logger setup
│   └── interfaces/
│       ├── dto/              # Request/response types
│       └── http/             # HTTP handlers, middleware, routes
└── deployments/docker/       # Dockerfile, docker-compose, nginx.conf
```
