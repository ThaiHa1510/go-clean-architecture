package database

import (
	"fmt"
 
	"github.com/Rayato159/go-clean-arch-v2/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
 )
 

type mysqlDatabase struct{
	db *gorm.Db
}

func(mysql *mysqlDatabase) GetDb(){
	return mysql.db
}

func NewMysqlDatabase(cfg *config.Config) Database{
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
  cfg.Db.Host,
  cfg.Db.User,
  cfg.Db.Password,
  cfg.Db.DBName,
  cfg.Db.Port,
  cfg.Db.SSLMode,
  cfg.Db.TimeZone,
 )
 
 db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
 if err != nil{
	panic("flailed to connect Database")
 }
 return &mysqlDatabase{db:db}
}