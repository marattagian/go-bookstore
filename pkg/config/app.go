package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
  database *gorm.DB 
  DSN = "host=localhost user=gian password=marattagian dbname=gorm port=5432"
)

func Connect()  {
  d, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
  if err != nil {
  	panic(err)
  }
  database = d
}

func GetDB() *gorm.DB {
	return database
}
