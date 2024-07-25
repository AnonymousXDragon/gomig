package main

import (
	"flag"
	"fmt"
	"gomig/internal/database"
	"gomig/internal/migration"
	"log"
)

func main() {
	dbConnection := flag.String("db", "", "Database connection string")
	migrationDir := flag.String("dir", "./migrations", "Databse Migration Dir")

	upFlag := flag.Bool("up", false, "Run Up Migrations")
	downFlag := flag.Bool("down", false, "Run Down Migrations")

	flag.Parse()

	if *dbConnection == "" {
		log.Fatal("please provide url to make connection with database")
	}

	db, err := database.Connect(*dbConnection)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	migrator, err := migration.NewMigrater(db, *migrationDir)
	fmt.Printf("%+v \n", migrator)
	if err != nil {
		log.Fatal(err.Error())
	}

	if *upFlag {
		if err := migrator.Up(); err != nil {
			log.Fatal(err.Error())
		}
	} else if *downFlag {
		if err := migrator.Down(); err != nil {
			log.Fatal(err.Error())
		}
	} else {
		log.Fatal("please specify up or down ...")
	}

	fmt.Println("Migrations successfully completed")
}
