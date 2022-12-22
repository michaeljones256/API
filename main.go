package main

import (
	"WebApp/api"
	"WebApp/database"
	"fmt"
)

func main() {
	//db := database.SetupDB()
	// database.SelectAllTable(db)

	db, err := database.Connect()
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	server := api.NewServer(db)
	server.StartServer()

}
