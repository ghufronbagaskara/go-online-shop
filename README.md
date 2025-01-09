# Online Shop Project

1. Run docker for PostgreSQL
```
docker run --name postgresql -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=database -d -p 5432:5432 postgres:16
```

2. Export environment variable
```
export DB_URI=postgres://user:password@localhost:5432/database?sslmode=disable
export ADMIN_SECRET=secret
```

3. Execute project
```
go run main.go
```



modules [
    1. github.com/gin-gonic/gin
    2. github.com/jackc/pgx
    3. github.com/google/uuid
    4. pkg.go.dev/golang.org/x/crypto
]
