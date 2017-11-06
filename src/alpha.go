package main

import (
	"database/sql"

	"github.com/MuchChaca/Training/simple-contact-list/src/handlers"
	"github.com/MuchChaca/Training/simple-contact-list/src/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db := connectDB("contactList")
	defer db.Close()

	models.Migrate(db)

	// New instance of echo
	e := echo.New()

	// g := e.Group("/index")

	// Serve the website
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../public",
	}))

	e.GET("/index", handlers.FindAll(db))
	e.POST("/add", handlers.AddContact(db))
	e.PUT("/mod", handlers.ModContact(db))
	e.DELETE("del/:id", handlers.DelContact(db))

	e.Logger.Fatal(e.Start(":8080"))
}

// initialise the db connecction
func connectDB(filepath string) *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(172.19.0.2:3306)/go_training_contactList") //? think we need the ip of the mysql container here

	// If any error on db then exit
	if err != nil {
		panic(err)
	}

	// if no errors but still don't get a db connection
	if db == nil {
		panic("db is nil")
	}

	return db
}

// // Code fore the database
// 	CREATE DATABASE IF NOT EXISTS contactList
// 		CHARACTER SET = utf8
// 		COLLATE = utf8_unicode_ci`
