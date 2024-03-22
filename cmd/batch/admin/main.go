package main

import (
	"log"

	"github.com/game-core/gocrafter/batch/game/presentation/router"
	"github.com/game-core/gocrafter/configs/database"
)

func main() {
	if _, err := database.InitMysql(); err != nil {
		log.Fatalf("failed to database.InitMysql: %v", err)
	}

	router.Router()
}
