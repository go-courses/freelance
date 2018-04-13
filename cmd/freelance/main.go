package main

import (
	"fmt"
	"log"

	server "github.com/go-courses/freelance/server"
)

func main() {
	// Read config from system environment
	/*c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err, "could not get env conf parms")
	}*/

	// подключение для PostgreSQL or MySQL, расскоментить нужное
	//conn, err := db.NewMySQL(c)
	/*conn, err := db.NewPgSQL(c)
	if err != nil {
		fmt.Println(err, " could not create database connection")
	}*/

	// этот код для миграции БД, если хотите откатить назад изменить на MigrateDown()
	/*err = conn.MigrateUp()
	if err != nil {
		fmt.Println(err, " could not migrate")
	}*/

	// Здесь код запуска rest api сервера
	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)
	restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)

	// fire the gRPC server in a goroutine
	go func() {
		err := server.StartGRPCServer(grpcAddress)
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
