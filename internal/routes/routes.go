package routes

import (
	"github.com/AlexGolangDeveloper/MusicLibraryAPI/internal/song"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	//Маршруты для работы с песнями

	// GetAllSongs получает список песен
	// @Summary Получить список песен
	// @Description Возвращает список всех песен из базы данных
	// @Tags songs
	// @Accept json
	// @Produce json
	// @Success 200 {array} song.Song
	// @Router /songs [get]
	r.HandleFunc("/songs", song.GetAllSongs).Methods("GET")
	r.HandleFunc("/songs/{id:[0-9]+}", song.GetSongByID).Methods("GET")
	r.HandleFunc("/songs/{id:[0-9]+}", song.DeleteSong).Methods("DELETE")
	r.HandleFunc("/songs/{id:[0-9]+}", song.UpdateSong).Methods("PUT")
	r.HandleFunc("/songs", song.AddSong).Methods("POST")

	// Подключение Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}

//1
