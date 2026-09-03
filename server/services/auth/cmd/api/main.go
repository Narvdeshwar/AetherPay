package main

import (
	"fmt"

	"github.com/Narvdeshwar/AetherPay/services/auth/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	db := config.InitDB(cfg)
	fmt.Println(db)
}
