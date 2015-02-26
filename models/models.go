package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var dbName = "./gonote.db"

type Category struct {
	Id   int64
	Name string `sql:"size:255"`
}

type Note struct {
	Id         int64
	CategoryID int64
	Note       string `sql:"size:1024"`
}

type Database struct {
	DB gorm.DB
}

func (d *Database) Connect() error {
	var err error
	d.DB, err = gorm.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) Setup() error {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		if err := d.Connect(); err != nil {
			return err
		}

		d.DB.DB()
		d.DB.CreateTable(&Category{})
		d.DB.CreateTable(&Note{})
	}
	return nil
}

func (d *Database) Close() {
	d.DB.Close()
}
