package main

import (
	handlers "github.com/karlygrcm/proxy-app/api/handlers"
	server "github.com/karlygrcm/proxy-app/api/server"
	utils "github.com/karlygrcm/proxy-app/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}
