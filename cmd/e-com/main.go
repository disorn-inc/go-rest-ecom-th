package main

import (
	"fmt"
	"os"

	"github.com/disorn-inc/go-rest-ecom-th/config"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	}
	return os.Args[1]
}

func main() {
	cfg := config.LoadConfig(envPath())
	fmt.Println("test")
	fmt.Println(cfg.App())
}