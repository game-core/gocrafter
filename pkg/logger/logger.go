package logger

import (
	"os"
)

// GenerateAppAdminDebug ログファイルを生成
func GenerateAppAdminDebug() (fp *os.File) {
	fp, err := os.OpenFile("./pkg/logger/api/admin/admin_debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp
}

// GenerateAppGameDebug ログファイルを生成
func GenerateAppGameDebug() (fp *os.File) {
	fp, err := os.OpenFile("./pkg/logger/api/game/game_debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp
}

// GenerateBatchGameDebug ログファイルを生成
func GenerateBatchGameDebug() (fp *os.File) {
	fp, err := os.OpenFile("./pkg/logger/batch/game/game_debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp
}

// GenerateBatchAdminDebug ログファイルを生成
func GenerateBatchAdminDebug() (fp *os.File) {
	fp, err := os.OpenFile("./pkg/logger/batch/admin/admin_debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp
}
