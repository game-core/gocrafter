package logger

import (
	"fmt"
	"os"
)

// DebugLog ログファイルを生成
func DebugLog(path, name string) (fp *os.File) {
	switch os.Getenv("APP_ENV") {
	case "prod":
		fp, err := prod(path, name)
		if err != nil {
			panic(err)
		}
		return fp
	case "dev":
		fp, err := dev(path, name)
		if err != nil {
			panic(err)
		}
		return fp
	}

	return fp
}

// prod 本番環境
func prod(path, name string) (*os.File, error) {
	return nil, nil
}

// dev dev環境
func dev(path, name string) (*os.File, error) {
	fp, err := os.OpenFile(fmt.Sprintf("%s/%s.log", path, name), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp, nil
}
