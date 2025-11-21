package router

import (
	"net/http"

	"github.com/tom150407/site-web-spot-api/controller"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Menu)
	mux.HandleFunc("/album/damso", controller.Damso)
	mux.HandleFunc("/track/laylow", controller.Laylow)

	return mux
}
