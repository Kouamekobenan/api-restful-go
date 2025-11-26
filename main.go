package main

import (
	"github.com/gin-gonic/gin"
	"go-user-api/database"
	"go-user-api/routes"

	// Swagger
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-user-api/docs" // chemin vers le package généré par swag init
)

/*
@title Go User API
@version 1.0
@description API pour gérer les utilisateurs (CRUD)
@termsOfService http://example.com/terms/

@contact.name API Support
@contact.url http://example.com
@contact.email support@example.com

@license.name MIT
@license.url https://opensource.org/licenses/MIT

@host localhost:8080
@BasePath /api/v1
*/
func main() {
	router := gin.Default()

	// Connexion à la base de données
	database.ConnectDatabase()

	// Enregistrer les routes utilisateurs
	routes.RegisterUserRoutes(router)
	// Route Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Lancer le serveur
	router.Run(":8080")
}
