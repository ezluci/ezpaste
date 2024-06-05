package main

import (
	"net/http"
	"os"
	"path"
	"net/url"
	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", mainPage)
	router.Get("/upload", uploadPaste)
	// router.Get("/admin", admin)
	router.Get("/{pasteId}", getPaste)


	http.ListenAndServe(":4000", router)
}


func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path.Join("html", "index.html"));
}


func getPaste(w http.ResponseWriter, r *http.Request) {

	pasteId := chi.URLParam(r, "pasteId")
	content, err := os.ReadFile(path.Join("pastes", pasteId))
	
	if err != nil {
		w.Header().Set("content-type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<span style=\"color: red\">error 404 not found :(</span>"))
		return
	}

	w.Write(content)
}


func generatePasteName() string {
	return genRandomString(5);
}


func uploadPaste(w http.ResponseWriter, r *http.Request) {

	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil || len(m["text"]) == 0 || len(m["text"][0]) == 0 {
		http.Redirect(w, r, "", http.StatusBadRequest)
		return
	}

	text := []byte(m["text"][0])

	var name string

	for name = generatePasteName(); ; name = generatePasteName() {
		// todo restrict "upload" from name
		_, err := os.ReadFile(path.Join("pastes", name))
	
		if err != nil {
			break
		}
	}
	
	// create folder 'pastes'
	_, err = os.Stat("pastes")
	if err != nil {
		os.Mkdir("pastes", 0750);
	}

	err = os.WriteFile(path.Join("pastes", name), text, 0644)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, string(name), http.StatusFound)
}