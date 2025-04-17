package main

import (
	"my-ops-admin/initalize"
)

func main() {
	router := initalize.Routers()
	router.Run("localhost:8080")
}
