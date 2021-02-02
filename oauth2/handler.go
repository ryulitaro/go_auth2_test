package oauth2

import (
	"log"
	"net/http"
	"os"

	"github.com/go-session/session"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loginHandler", r.Method)

	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		log.Println("loginHandler", r.Form)

		if r.Form == nil {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		store.Set("LoggedInUserID", r.Form.Get("username"))
		store.Save()

		log.Println("loginHandler", r.Form.Get("username"), r.Form.Get("password"))

		w.Header().Set("Location", "/auth")
		w.WriteHeader(http.StatusFound)
		return
	}
	outputHTML(w, r, "static/login.html")
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("authHandler")

	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		log.Println(store.Get("LoggedInUserID"))
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	outputHTML(w, r, "static/auth.html")
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	log.Println("outputHTML", filename, req.Header)
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
	log.Println("outputHTML", w.Header())
}
