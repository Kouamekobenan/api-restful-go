package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"go-user-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// On construit un chemin absolu pour éviter les problèmes relatifs
	execDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("❌ Impossible de récupérer le répertoire courant : %v", err)
	}

	dbPath := filepath.Join(execDir, "users.db")

	// Vérifie si le fichier existe, sinon le crée
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatalf("❌ Impossible de créer le fichier SQLite : %v", err)
		}
		file.Close()
		fmt.Println("🆕 Fichier users.db créé.")
	}

	// Connexion à SQLite
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Erreur de connexion à la base de données SQLite : %v", err)
	}

	// Migration du modèle
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("❌ Erreur lors de la migration : %v", err)
	}

	DB = db
	fmt.Println("✅ Base de données connectée avec succès à :", dbPath)
}
