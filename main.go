package main

import (
	"WebApp/api"
	"WebApp/database"
	"fmt"
)

func main() {
	//db := database.SetupDB()
	// database.SelectAllTable(db)

	database, err := database.Connect()
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	server := api.NewServer(database)
	server.StartServer()

}
