package main

import (
	"fmt"
	"github.com/AlexGolangDeveloper/MusicLibraryAPI/config"
	_ "github.com/AlexGolangDeveloper/MusicLibraryAPI/docs" // Swagger-документация
	"github.com/AlexGolangDeveloper/MusicLibraryAPI/internal/database"
	"github.com/AlexGolangDeveloper/MusicLibraryAPI/internal/routes"
	"github.com/AlexGolangDeveloper/MusicLibraryAPI/internal/song"
	"log"
	"net/http"
)

// @title Music Library API
// @version 1.0
// @description API для управления музыкальной библиотекой
// @host localhost:8080
// @BasePath /

func main() {

	cfg := config.LoadConfig()

	//Подключение к базе данных
	database.ConnectDatabase(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	//Миграция таблицы песен
	err := database.DB.AutoMigrate(&song.Song{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	//Настраиваем маршруты
	r := routes.RegisterRoutes()

	//Запускаем сервер
	fmt.Printf("Server running on port %s\n", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}

//1
