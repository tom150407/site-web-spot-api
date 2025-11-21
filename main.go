package main

import (
	"log"
	"net/http"

	"github.com/tom150407/site-web-spot-api/router"
)

func main() {

	r := router.NewRouter()

	// fichiers statiques
	fs := http.FileServer(http.Dir("asset"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))

	http.Handle("/", r)

	log.Println("🔥 Serveur prêt : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
