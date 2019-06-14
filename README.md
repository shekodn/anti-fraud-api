# Anti Fraud API

### How To Run

1. Clone the project
2. Run ``` glide install ```
3. Set up an .env file in the project's root directory
```
POSTGRES_DB=anti-fraud-api
POSTGRES_PASSWORD=mysecretdbpassword
POSTGRES_USER=postgres
DB_TYPE=postgres
DB_HOST=localhost
DB_HOST_DOCKER=host.docker.internal
DB_PORT=5432
CACHE_PORT=6379
CACHE_PASSWORD=mysecretcachepassword
PORT=8000

```
4. Simulate an environment with Redis (cache), Postgres (db), and the
Application by running ``` docker-compose up --build ```

