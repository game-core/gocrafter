package log

import (
	"os"
)

func GenerateAdminLog() (fp *os.File) {
	fp, err := os.OpenFile("log/admin_debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp
}

func GenerateAuthLog() (fp *os.File) {
	fp, err := os.OpenFile("log/auth_debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp
}

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
