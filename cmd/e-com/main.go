package main

import (
	"fmt"
	"os"

	"github.com/disorn-inc/go-rest-ecom-th/config"
	"github.com/disorn-inc/go-rest-ecom-th/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	}
	return os.Args[1]
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	fmt.Println(db)
}