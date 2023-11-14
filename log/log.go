package log

import (
	"os"
)

func GenerateApiLog() (fp *os.File) {
	fp, err := os.OpenFile("log/api_debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp
}

func GenerateBatchLog() (fp *os.File) {
	fp, err := os.OpenFile("log/batch_debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp
}
