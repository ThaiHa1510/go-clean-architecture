package database 

import "gorm.io/gorm"

type Database struct{
	GetDb() *gorm.Db
}