# by_te : Simplifier

A simplifier service built in Go, with MySQL for persistence, Redis for caching, and Nginx for load balancing across two app instances.

---

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.22+](https://go.dev/dl/) _(only needed for local development)_

---

## Installation & Setup

### Option 1: Docker (recommended)

1. **Clone the repository**

```bash
   git clone 
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
   | MySQL | localhost:3306 |
   | Redis | localhost:6379 |

   > Services start in dependency order: MySQL and Redis → migrations → app instances → Nginx.
   > Nginx will only start once both app instances pass their healthchecks.

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
   docker compose -f deployments/docker/docker-compose.yml down -v --remove-orphans
```

---

### Option 2: Local Development

1. **Clone the repository**

```bash
   git clone 
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
   DB_PORT=3306
   REDIS_ADDR=localhost:6379
```

5. **Run migrations**

```bash
   docker compose -f deployments/docker/docker-compose.yml up migrate
```

6. **Run the application**

```bash
   go run ./cmd/api
```

7. **Verify**

```bash
   curl http://localhost:8080/health
   # OK
```

---

## Fresh Start (clean rebuild)

To wipe everything and start from scratch:

```bash
docker compose -f deployments/docker/docker-compose.yml down -v --remove-orphans
docker compose -f deployments/docker/docker-compose.yml up --build
```

To watch startup logs in order:

```bash
docker compose -f deployments/docker/docker-compose.yml logs -f migrate app-1 app-2 nginx
```

---

## Debugging

**Check service health status**
```bash
docker inspect by_te_mysql | grep -A 10 "Health"
```

**Check logs with timestamps**
```bash
docker compose -f deployments/docker/docker-compose.yml logs --timestamps mysql migrate
```

**Manually test DB connectivity**
```bash
docker run --rm --network by_te_app-network mysql:8 \
  mysqladmin ping -h mysql -uroot -proot
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
├── migrations/               # SQL migration files
└── deployments/docker/       # Dockerfile, docker-compose, nginx.conf
```