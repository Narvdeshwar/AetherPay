package main

import (
	"fmt"

	"github.com/Narvdeshwar/AetherPay/shared"
)

func main() {
	tenantID := "test_user_101"
	err := shared.NewAPIError(401, "Invalid JWT Token", tenantID)
	fmt.Println("Auth service running")
	fmt.Println(err.Error())
}
