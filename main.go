package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
) 

	

func main() {
	db, err := sql.Open("pgx", os.Getenv("DB_URI")) // using env variable at terminal : export DB_URI=postgres://user:password@localhost:5432/database?sslmode=disable
	if err != nil {
		fmt.Printf("Error connect to db: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()


	if err = db.Ping(); err != nil {
		fmt.Printf("Error to verif db connection: %v\n", err)
		os.Exit(1)
	}
}