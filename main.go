package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

type Film struct {
	Title    string
	Director string
}

var allFilms = []Film{
	{Title: "The Godfather", Director: "Francis Ford Coppola"},
	{Title: "Blade Runner", Director: "Ridley Scott"},
	{Title: "The Thing", Director: "John Carpenter"},
}

func main() {
	fmt.Println("Go app...")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": allFilms,
		}
		tmpl.Execute(w, films)
	}

	// handler function #2 - returns the template block with the newly added film, as an HTMX response
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		newFilm := Film{Title: title, Director: director}
		allFilms = append(allFilms, newFilm)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", newFilm)
	}

	// handler function #3 - handles filtering of films
	h3 := func(w http.ResponseWriter, r *http.Request) {
		filter := r.URL.Query().Get("filter")
		filteredFilms := []Film{}

		for _, film := range allFilms {
			if strings.Contains(strings.ToLower(film.Title), strings.ToLower(filter)) ||
				strings.Contains(strings.ToLower(film.Director), strings.ToLower(filter)) {
				filteredFilms = append(filteredFilms, film)
			}
		}

		tmpl := template.Must(template.New("film-list").Parse(`
			{{ range . }}
			<li class="list-group-item bg-primary text-white">{{ .Title }} - {{ .Director }}</li>
			{{ else }}
			<li class="list-group-item">No films found.</li>
			{{ end }}
		`))

		tmpl.Execute(w, filteredFilms)
	}

	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
	http.HandleFunc("/filter-films", h3)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
