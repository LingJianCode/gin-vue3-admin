package main

import (
	"my-ops-admin/initalize"
)

func main() {
	initalize.InitViper()
	initalize.InitZap()
	initalize.InitDB()
	initalize.AutoMigrate()
	router := initalize.Routers()
	router.Run("localhost:8080")
}
