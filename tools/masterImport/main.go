package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func main() {
	if err := MasterImport(); err != nil {
		fmt.Println(err)
	}
}

func MasterImport() error {
	token := os.Getenv("GITHUB_TOKEN")
	organization := os.Getenv("GITHUB_NAME")
	repo := os.Getenv("GITHUB_REPOSITORY_NAME")
	path := os.Getenv("GITHUB_REPOSITORY_PATH")

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("MASTER_MYSQL_USER"),
			os.Getenv("MASTER_MYSQL_PASSWORD"),
			os.Getenv("MASTER_MYSQL_WRITE_HOST"),
			os.Getenv("MASTER_MYSQL_DATABASE"),
		),
	)
	if err != nil {
		return err
	}
	defer db.Close()

	client := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})))
	_, contents, _, err := client.Repositories.GetContents(context.Background(), organization, repo, path, nil)
	if err != nil {
		return err
	}

	for _, content := range contents {
		if content.Type != nil && *content.Type == "file" && strings.HasSuffix(*content.Name, ".csv") {
			fileContent, _, err := client.Repositories.DownloadContents(context.Background(), organization, repo, *content.Path, nil)
			if err != nil {
				return err
			}
			defer fileContent.Close()

			if err := bulkInsert(db, csv.NewReader(fileContent), *content.Name); err != nil {
				return err
			}

			fmt.Println(fmt.Sprintf("Import: %s", *content.Name))
		}
	}

	return nil
}

func bulkInsert(db *sql.DB, reader *csv.Reader, name string) error {
	columns, err := reader.Read()
	if err != nil {
		return err
	}

	var valueArgs []interface{}
	var valuePlaceholders []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		var recordValues []string
		for _, v := range record {
			recordValues = append(recordValues, "?")
			if v != "" {
				valueArgs = append(valueArgs, v)
			} else {
				valueArgs = append(valueArgs, nil)
			}
		}
		valuePlaceholders = append(valuePlaceholders, "("+strings.Join(recordValues, ",")+")")
	}

	if len(valuePlaceholders) <= 0 {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			if err := tx.Commit(); err != nil {
				return
			}
		} else {
			if err := tx.Rollback(); err != nil {
				return
			}
		}
	}()

	if _, err := tx.Exec(fmt.Sprintf("DELETE FROM `%s`", strings.ReplaceAll(name, ".csv", ""))); err != nil {
		return err
	}

	if _, err := tx.Exec(fmt.Sprintf("INSERT INTO `%s` (`%s`) VALUES %s", strings.ReplaceAll(name, ".csv", ""), strings.Join(columns, "`,`"), strings.Join(valuePlaceholders, ",")), valueArgs...); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
