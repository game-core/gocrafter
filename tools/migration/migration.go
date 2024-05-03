package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := migrateCommonDB("./docs/sql/common/ddl"); err != nil {
		fmt.Println("error migrating the database:", err)
		return
	}

	if err := migrateMasterDB("./docs/sql/master/ddl"); err != nil {
		fmt.Println("error migrating the database:", err)
		return
	}

	if err := migrateUserDB("./docs/sql/user/ddl"); err != nil {
		fmt.Println("error migrating the database:", err)
		return
	}

	fmt.Println("migration completed successfully.")
}

func migrateCommonDB(migrationsDir string) error {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("COMMON_MYSQL_USER"),
		os.Getenv("COMMON_MYSQL_PASSWORD"),
		os.Getenv("COMMON_MYSQL_WRITE_HOST"),
		os.Getenv("COMMON_MYSQL_DATABASE"),
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	defer db.Close()

	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			filePath := filepath.Join(migrationsDir, file.Name())
			query, err := ioutil.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("error reading SQL file %s: %v", file.Name(), err)
			}

			tx, err := db.Begin()
			if err != nil {
				return err
			}

			if _, err := tx.Exec(string(query)); err != nil {
				tx.Rollback()
				return fmt.Errorf("error executing SQL file %s: %v", file.Name(), err)
			}

			if err := tx.Commit(); err != nil {
				tx.Rollback()
				return err
			}

			fmt.Printf("executed migration %s\n", file.Name())
		}
	}

	return nil
}

func migrateMasterDB(migrationsDir string) error {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MASTER_MYSQL_USER"),
		os.Getenv("MASTER_MYSQL_PASSWORD"),
		os.Getenv("MASTER_MYSQL_WRITE_HOST"),
		os.Getenv("MASTER_MYSQL_DATABASE"),
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	defer db.Close()

	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			filePath := filepath.Join(migrationsDir, file.Name())
			query, err := ioutil.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("error reading SQL file %s: %v", file.Name(), err)
			}

			tx, err := db.Begin()
			if err != nil {
				return err
			}

			if _, err := tx.Exec(string(query)); err != nil {
				tx.Rollback()
				return fmt.Errorf("error executing SQL file %s: %v", file.Name(), err)
			}

			if err := tx.Commit(); err != nil {
				tx.Rollback()
				return err
			}

			fmt.Printf("executed migration %s\n", file.Name())
		}
	}

	return nil
}

func migrateUserDB(migrationsDir string) error {
	shardCountStr := os.Getenv("MYSQL_SHARD_COUNT")
	shardCount, err := strconv.Atoi(shardCountStr)
	if err != nil {
		return err
	}

	for i := 0; i <= shardCount; i++ {
		shard := fmt.Sprintf("_%v", i)
		dataSourceName := fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv(fmt.Sprintf("USER_MYSQL_USER%s", shard)),
			os.Getenv(fmt.Sprintf("USER_MYSQL_PASSWORD%s", shard)),
			os.Getenv(fmt.Sprintf("USER_MYSQL_WRITE_HOST%s", shard)),
			os.Getenv(fmt.Sprintf("USER_MYSQL_DATABASE%s", shard)),
		)

		db, err := sql.Open("mysql", dataSourceName)
		if err != nil {
			return err
		}
		defer db.Close()

		files, err := ioutil.ReadDir(migrationsDir)
		if err != nil {
			return err
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".sql") {
				filePath := filepath.Join(migrationsDir, file.Name())
				query, err := ioutil.ReadFile(filePath)
				if err != nil {
					return fmt.Errorf("error reading SQL file %s: %v", file.Name(), err)
				}

				tx, err := db.Begin()
				if err != nil {
					return err
				}

				if _, err := tx.Exec(string(query)); err != nil {
					tx.Rollback()
					return fmt.Errorf("error executing SQL file %s: %v", file.Name(), err)
				}

				if err := tx.Commit(); err != nil {
					tx.Rollback()
					return err
				}

				fmt.Printf("executed migration %s\n", file.Name())
			}
		}
	}

	return nil
}
