package db

import (
	"fmt"

	"github.com/eigakan/user-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(conf config.DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
