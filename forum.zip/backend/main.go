package main

import (
	"fmt"
	"forum/config"
	"forum/routes"
	"net/http"
	"os"
)

func main() {
	config.DB = config.InitDB("../database/forum.db")
	config.CreateDatabaseTables(config.DB, "../database/schema.sql")
	defer config.DB.Close()
	address := "localhost:8080"
	// home routes
	routes.HomeRoute()
	// authentication routes
	routes.AuthRoutes()
	// post routes
	routes.PostRoute(config.DB)
	//users routes
	// fmt.Println("enter to thsi step . ")
	//
	routes.Usersapi()
	// like routes
	routes.ReactionsRoute(config.DB)
	// comment routes
	routes.CommentsRoute(config.DB)
	//categoeies routes
	routes.CategoriesRoute(config.DB)
	//websocket ... 
	routes.WebSocket()
	
	routes.FilterRoute(config.DB)
	fmt.Printf("Server is running on http://%s \n", address)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		os.Exit(1)
	}
}
