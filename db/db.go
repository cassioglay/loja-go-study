package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_user := os.Getenv("POSTGRES_USER")
	db_password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	db_host := os.Getenv("DB_HOST")

	conexao := "user=" + db_user + " dbname=" + database + " password=" + db_password + " host=" + db_host + " sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
