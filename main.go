package main

import (
	"CNLib-backend/module"
	"CNLib-backend/utility"
)

func main() {
	server := utility.NewMyServer(true)
	server.AddRoute("/login", module.Login)
	server.AddRoute("/logout", module.Logout)
	server.AddRoute("/test", module.Test)
	server.Run()
}
