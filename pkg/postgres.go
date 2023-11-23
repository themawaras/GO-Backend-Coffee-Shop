package pkg

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func PostgreSQLDB() (*sqlx.DB, error) {
	// myEnv, _ := godotenv.Read()

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
	return sqlx.Connect("postgres", config)
}
