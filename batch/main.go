package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/game-core/gocrafter/batch/di"
	batchLog "github.com/game-core/gocrafter/log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// di: wire ./batch/di/wire.go
	exampleCommand := di.InitializeExampleCommand()

	// Batch list
	batchList := map[string]func() error{
		"ListExample": exampleCommand.ListExample,
	}

	commandPtr := flag.String("command", "", "Specify the command")
	flag.Parse()
	command := *commandPtr
	selectedCommand := command

	// Run batch
	if batch, ok := batchList[selectedCommand]; ok {
		logFile := batchLog.GenerateBatchLog()
		log.SetOutput(logFile)

		err := batch()
		if err != nil {
			log.Println(err)
		}
	} else {
		fmt.Println("Invalid command")
	}
}
