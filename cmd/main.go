package main

import (
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/sellto/myawesomecatapp/assets"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index").Parse(assets.Page)
		//tmpl, err := template.New("index").Parse(assets.Page_v2)
		if err != nil {
			http.Error(w, "Failed to parse template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/randomcat", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
		if err != nil {
			http.Error(w, "Failed to get random cat", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	})

	http.HandleFunc("/catimage", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to get cat image", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		io.Copy(w, resp.Body)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
