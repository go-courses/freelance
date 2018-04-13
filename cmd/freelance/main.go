package main

import (
	"fmt"
	"log"

	"github.com/go-courses/freelance/config"
	"github.com/go-courses/freelance/db"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err, "could not get env conf parms")
	}

	conn, err := db.NewMySQL(c)
	if err != nil {
		fmt.Println(err, " could not create database connection")
	}

	conn, err := db.NewPgSQL(c)
	if err != nil {
		fmt.Println(err, " could not create database connection")
	}

	err = conn.MigrateUp()
	if err != nil {
		//spew.Dump(err)
		fmt.Println(err, " could not migrate")
	}
}
