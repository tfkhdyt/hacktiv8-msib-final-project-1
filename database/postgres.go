package database

import (
	"hacktiv8-msib-final-project-1/config"
	"log"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(config.GetDBConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Fatalln(db.AutoMigrate())

	log.Println("Connected to DB!")
}

func GetPostgresInstance() *gorm.DB {
	return db
}
