package handlers

import (
	"github.com/kataras/iris"
)

// HandlerRedirection should redirect traffict app
func HandlerRedirection(app *iris.Application) {
	app.Get("/ping", func(context iris.Context) {
		context.JSON(iris.Map{"result": "ok"})
	})
}
