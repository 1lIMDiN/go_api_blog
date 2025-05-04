package main

import (
	"fmt"

	"apiBlog/pkg/db"
	"apiBlog/pkg/server"
)

func main() {
	dbFile := "blog.db"

	if err := db.Init(dbFile); err != nil {
		fmt.Errorf("Failed to initialize database: %w", err)
	}
	defer db.DB.Close()

	port := 7540
	if err := server.Run(port); err != nil {
		fmt.Errorf("Failed to start server: %w", err)
	}
}