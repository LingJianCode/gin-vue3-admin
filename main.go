package main

import (
	"my-ops-admin/initalize"
)

func main() {
	initalize.InitViper()
	initalize.InitZap()
	initalize.InitDB()
	initalize.AutoMigrate()
	initalize.InitValidator()
	router := initalize.Routers()
	router.Run("localhost:8080")
}
