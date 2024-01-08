package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&ExampleModel{})
	fmt.Println("Migration successful")

	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("Hello World every 10 seconds")
		task(db)
	})

	c.Start()

	select {}
}
