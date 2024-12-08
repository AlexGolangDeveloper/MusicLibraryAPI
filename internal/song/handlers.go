package song

import (
	"encoding/json"
	"github.com/AlexGolangDeveloper/MusicLibraryAPI/internal/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// GetAllSongs возвращает все песни
func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	var songs []Song
	database.DB.Find(&songs)
	err := json.NewEncoder(w).Encode(songs)
	if err != nil {
		return
	}
}

//1

// AddSong добавляет новую песню
func AddSong(w http.ResponseWriter, r *http.Request) {
	var song Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	database.DB.Create(&song)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(song)
	if err != nil {
		return
	}

}

// GetSongByID возвращает песню по id
func GetSongByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Id is not correct", http.StatusNotFound)
		return
	}
	var song Song
	if err := database.DB.First(&song, id).Error; err != nil {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(song)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// UpdateSong обновляет данные песни
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Id is not correct", http.StatusNotFound)
		return
	}
	var song Song
	if err := database.DB.First(&song, id); err != nil {
		http.Error(w, "Invalid not found", http.StatusBadRequest)
		return
	}
	var updatedSong Song
	if err := json.NewDecoder(r.Body).Decode(&updatedSong); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	song.Group = updatedSong.Group
	song.Song = updatedSong.Song
}

// DeleteSong удаляет песню
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Id is not correct", http.StatusNotFound)
		return
	}
	if err := database.DB.Delete(&Song{}, id).Error; err != nil {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
