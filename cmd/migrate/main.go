package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/ramdhanrizkij/ara-store/internal/config"
)

func main() {
	cfg := config.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide migration command: up | down")
	}

	cmd := os.Args[1]

	switch cmd {
	case "up":
		if err := goose.Up(db, "migrations"); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unknown command. Use: up | down")
	}

	fmt.Println("Migration executed:", cmd)
}
