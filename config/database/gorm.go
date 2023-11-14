package database

import (
	"fmt"
	"os"
	
	"github.com/jinzhu/gorm"
)

type SqlHandler struct {
	User   *Conn
	Master *Conn
}

type Conn struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
}

func NewDB() *SqlHandler {
    return &SqlHandler{
		User:   userDB(),
		Master: masterDB(),
    }
}

func userDB() *Conn {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_READ_USER"),
		os.Getenv("MYSQL_READ_PASSWORD"),
		os.Getenv("MYSQL_READ_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_WRITE_USER"),
		os.Getenv("MYSQL_WRITE_PASSWORD"),
		os.Getenv("MYSQL_WRITE_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

    readDB, err := gorm.Open("mysql", readConn)
    if err != nil {
        panic(err.Error())
    }

    writeDB, err := gorm.Open("mysql", writeConn)
    if err != nil {
        panic(err.Error())
    }

	readDB.SingularTable(true)
	writeDB.SingularTable(true)

	return &Conn{
		ReadConn:  readDB,
        WriteConn: writeDB,
	}
}

func masterDB() *Conn {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_READ_USER"),
		os.Getenv("MYSQL_READ_PASSWORD"),
		os.Getenv("MYSQL_READ_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_WRITE_USER"),
		os.Getenv("MYSQL_WRITE_PASSWORD"),
		os.Getenv("MYSQL_WRITE_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	
    readDB, err := gorm.Open("mysql", readConn)
    if err != nil {
        panic(err.Error())
    }

    writeDB, err := gorm.Open("mysql", writeConn)
    if err != nil {
        panic(err.Error())
    }

	readDB.SingularTable(true)
	writeDB.SingularTable(true)

	return &Conn{
		ReadConn:  readDB,
        WriteConn: writeDB,
	}
}
