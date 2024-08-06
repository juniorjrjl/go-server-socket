package main

import (
	"os"
	"os/exec"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbHost == "" || dbPort == "" || dbDatabase == "" || dbUser == "" || dbPassword == "" {
		panic("Uma ou mais variáveis de ambiente estão ausentes")
	}
	/*if err := godotenv.Load(); err != nil {
		panic(err)
	}*/

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
