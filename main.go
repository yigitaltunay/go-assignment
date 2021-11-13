// Yigit Altunay
// Go Assignment Project

package main

import (
	"github.com/yigitaltunay/go-assignment/config"
	"github.com/yigitaltunay/go-assignment/routes"
)

// TODO Swagger
func init() { config.InitServe() }

func main() { routes.InitializeRoutes() }
