package config

import (
        "gorm.io/driver/mysql"
        "gorm.io/gorm"
)

var (
    db  *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:769321djR@tcp(127.0.0.1:3306)/bookmgmnt?charset=utf8mb4&parseTime=True&loc=Local")
}

func GetDB() *gorm.DB {
        return db
}