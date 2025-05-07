package main

import (
	"my-ops-admin/initalize"
	mycasbin "my-ops-admin/pkg/my_casbin"
)

func main() {
	initalize.InitViper()
	initalize.InitZap()
	initalize.InitDB()
	initalize.AutoMigrate()
	initalize.InitValidator()
	mycasbin.Casbin.InitCasbin()
	router := initalize.Routers()
	router.Run("localhost:8080")
}
