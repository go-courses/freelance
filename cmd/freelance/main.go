package main

import (
	"fmt"
	"log"

	"github.com/go-courses/freelance/api"
	"github.com/go-courses/freelance/config"
	"github.com/go-courses/freelance/db"
	server "github.com/go-courses/freelance/server"
)

func main() {
	// Read config from system environment
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err, "could not get env conf parms")
	}

	s, err := api.NewServer(c)

	// для миграции БД, для отката использовать MigrateDown()
	if c.DoMigration == "Yes" {
		switch c.DbType {
		case "mysql":
			conn, err := db.NewMySQL(c)
			if err != nil {
				log.Fatalf("failed connect to Mysql: %s", err)
			}
			err = conn.MigrateUp(c.MigrationsFolder)
			if err != nil {
				log.Fatal(err)
			}
		case "postgres":
			conn, err := db.NewPgSQL(c)
			if err != nil {
				log.Fatalf("failed connect to PostgreSQL: %s", err)
			}
			err = conn.MigrateUp(c.MigrationsFolder)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Здесь код запуска rest api сервера
	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)
	restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)

	// fire the gRPC server in a goroutine
	go func() {
		err := server.StartGRPCServer(grpcAddress, s)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := server.StartRESTServer(restAddress, grpcAddress)
		if err != nil {
			log.Fatalf("failed to start REST server: %s", err)
		}
	}()

	// infinite loop
	log.Printf("Entering infinite loop")
	select {}

}
