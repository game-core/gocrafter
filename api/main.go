package main

import (
	"github.com/game-core/gocrafter/api/presentation/router"
)

// @title Golang
// @version 1.0
// @description This is a sample swagger server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8001
// @BasePath /
func main() {
	router.Init()
}
