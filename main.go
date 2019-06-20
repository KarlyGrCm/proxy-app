package main

import (
	handlers "github.com/karlygrcm/proxy-app/api/handlers"
	"github.com/karlygrcm/proxy-app/api/server"
	"github.com/karlygrcm/proxy-app/api/utils"
)

func main() {
	/*
		Router Iris
		Env vars
	*/
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}
