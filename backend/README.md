# Database Migrations

## Running Migrations

1. Ensure you have [golang-migrate](https://github.com/golang-migrate/migrate) installed.
2. Run the following command in the terminal:

```bash
migrate -database "${DB_URL}?sslmode=require&x-migrations-table-quoted=true&x-migrations-table=%22${MICROSERVICE}%22.%22schema_migrations%22" -path db/migrations up
```

3. To roll back the migrations, run the following command in the terminal:
```bash
migrate -database "${DB_URL}?sslmode=require&x-migrations-table-quoted=true&x-migrations-table=%22${MICROSERVICE}%22.%22schema_migrations%22" -path db/migrations down
```

> **_NOTE:_** `${DB_URL}` should be your database connection string. 
> `${MICROSERVICE}` is the schema (namespace) for your service, eg. donation_service, campaign_service etc.
>
> The `x-migrations-table` option scopes the migrations table to the microservice schema for development purposes.
> In production, each microservice should ideally have its own dedicated database for stronger isolation and maintainability.
