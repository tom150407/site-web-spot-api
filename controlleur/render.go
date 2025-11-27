package controlleur

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	path := filepath.Join("template", tmpl)

	log.Println("➡️ Render :", path)

	t, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return t.Execute(w, data)
}
