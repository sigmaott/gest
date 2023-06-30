package main

import (
	"usage-monnitor/src/module"

	_ "github.com/gestgo/gest/package/technique/version"
)

// @title Gest Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
func main() {
	app := module.NewApp()
	app.Run()
}
