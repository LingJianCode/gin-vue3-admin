package main

import (
	"my-ops-admin/initalize"
)

func main() {
	initalize.InitViper()
	initalize.InitZap()
	router := initalize.Routers()
	router.Run("localhost:8080")
}
