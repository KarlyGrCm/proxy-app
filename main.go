package main

import (
	handlers "github.com/karlygrcm/proxy-app/api/handlers"
	middleware "github.com/karlygrcm/proxy-app/api/middleware"
	"github.com/karlygrcm/proxy-app/api/server"
	"github.com/karlygrcm/proxy-app/api/utils"
)

// the middleware is going to pass the request to a queue
// the middleware organice the requests based on the criteria
//so the dispatcher is going to pop everything in order

func main() {
	/*
		Router Iris
		Env vars
	*/
	utils.LoadEnv()
	app := server.SetUp()
	middleware.InitQueue()
	handlers.HandlerRedirection(app)
	server.RunServer(app)

}
