package handlers

import (
	middleware "github.com/karlygrcm/proxy-app/api/middleware"
	"github.com/kataras/iris"
)

// HandlerRedirection should redirect traffic
func HandlerRedirection(app *iris.Application) {
	app.Get("/ping", middleware.ProxyMiddleware, proxyHandler)
}

func proxyHandler(context iris.Context) {
	context.JSON(iris.Map{"result": "ok"})
}
