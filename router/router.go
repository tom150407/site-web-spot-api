package router

import (
	"net/http"
	"site-web-spot/controlleur"
)

func NewRouter() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/", controlleur.Menu)
    mux.HandleFunc("/album/damso", controlleur.Damso)
    mux.HandleFunc("/track/laylow", controlleur.Laylow)

    return mux
}
