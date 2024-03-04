package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// GenerateLog ログファイルを生成する
func GenerateLog(path, name string) (fp *os.File) {
	switch os.Getenv("APP_ENV") {
	case "prod":
		fp, err := generateProd(path, name)
		if err != nil {
			panic(err)
		}
		return fp
	case "dev":
		fp, err := generateDev(path, name)
		if err != nil {
			panic(err)
		}
		return fp
	default:
		return nil
	}
}

// CreateLog ログを追加する
func CreateLog(l struct{}) {
	switch os.Getenv("APP_ENV") {
	case "prod":
		if err := createProd(l); err != nil {
			log.Fatal(err)
		}
	case "dev":
		if err := createDev(l); err != nil {
			log.Fatal(err)
		}
	default:
	}
}

// generateProd 本番環境のログファイルを生成する
func generateProd(path, name string) (*os.File, error) {
	return nil, nil
}

// generateDev dev環境のログファイルを生成する
func generateDev(path, name string) (*os.File, error) {
	fp, err := os.OpenFile(fmt.Sprintf("%s/%s.log", path, name), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp, nil
}

// createProd 本番環境のログを追加する
func createProd(l any) error {
	return nil
}

// createDev dev環境のログを追加する
func createDev(l any) error {
	j, err := json.Marshal(l)
	if err != nil {
		return err
	}

	log.Print(string(j))
	return nil
}
