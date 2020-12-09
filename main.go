package main

import (
	"github.com/marshhu/ma-api/router"
	"github.com/marshhu/ma-frame/app"
)
func main() {
    app := app.NewApp(8081)
    app.RegisterRouter(router.Register())
    app.Run()
}
