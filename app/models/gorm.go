package models

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var DB **gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", getConnectString())

	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}

	db.DB()
	DB = &db
}

type Model struct {
	gorm.Model
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteedAt *time.Time
}

type Validator interface {
	IsSatisfied(interface{}) bool
	DefaultMessage() string
}

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Could not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectString() string {
	host := getParamString("db.host", "localhost")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "root")
	pass := getParamString("db.password", "")
	dbname := getParamString("db.name", "studygo")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")
	timezone := getParamString("db.timezone", "parseTime=true&loc=Asia%2FTokyo")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s?%s", user, pass, protocol, host, port, dbname, dbargs, timezone)
}
