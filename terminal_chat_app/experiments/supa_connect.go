package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	// Replace with your Supabase connection string
	// postgresql://postgres:[YOUR-PASSWORD]@db.latavriesquutbnjvhdi.supabase.co:5432/postgres
	connString := "postgres://postgres:TCA@supa12345@db.latavriesquutbnjvhdi.supabase.co:5432/postgres?sslmode=require"

	// Create a connection to the database
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Example: Query data from a table called "profiles"
	rows, err := conn.Query(context.Background(), "SELECT * FROM profiles")
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var id int
		var name string
		// Adjust the variables according to your table structure
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatalf("Row scan failed: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v", err)
	}
}
