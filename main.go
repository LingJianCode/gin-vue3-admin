package main

import (
	"my-ops-admin/initialize"
	mycasbin "my-ops-admin/pkg/my_casbin"
)

func main() {
	initialize.InitViper()
	initialize.InitZap()
	initialize.InitDB()
	initialize.AutoMigrate()
	initialize.InitValidator()
	mycasbin.Casbin.InitCasbin()
	router := initialize.Routers()
	router.Run("localhost:8080")
}
