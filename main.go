package main

import (
    "log"
    "net/http"

    "github.com/joho/godotenv"
    "site_web_spot/router"
)

func main() {

    if err := godotenv.Load(); err != nil {
        log.Println("⚠️ Impossible de charger .env, vérifie sa position")
    }

    r := router.NewRouter()

    fs := http.FileServer(http.Dir("asset"))
    http.Handle("/asset/", http.StripPrefix("/asset/", fs))

    http.Handle("/", r)

    log.Println("🔥 Serveur prêt : http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
