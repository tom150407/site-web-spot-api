package controller

import (
	"log"
	"net/http"
)

func Menu(w http.ResponseWriter, r *http.Request) {
	log.Println("➡️ GET /")
	RenderTemplate(w, "menu.html", nil)
}

func Damso(w http.ResponseWriter, r *http.Request) {
	log.Println("➡️ GET /album/damso")

	// Appel API Spotify
	albums, err := GetDamsoAlbums()
	if err != nil {
		http.Error(w, "Erreur Spotify", 500)
		return
	}

	RenderTemplate(w, "album_damso.html", albums)
}

func Laylow(w http.ResponseWriter, r *http.Request) {
	log.Println("➡️ GET /track/laylow")

	track, err := GetLaylowTrack()
	if err != nil {
		http.Error(w, "Erreur Spotify", 500)
		return
	}

	RenderTemplate(w, "track_laylow.html", track)
}
