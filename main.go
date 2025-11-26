package main

import (
	"github.com/gin-gonic/gin"
	"go-user-api/database"
	"go-user-api/routes"
)

func main() {
	router := gin.Default()
	// Connexion à la base
	database.ConnectDatabase()

	// Enregistrer les routes
	routes.RegisterUserRoutes(router)

	// Lancer le serveur
	router.Run(":8080")
}
