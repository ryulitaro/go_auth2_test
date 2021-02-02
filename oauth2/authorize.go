package oauth2

import (
	"log"
	"net/http"
	"net/url"

	"github.com/go-session/session"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	log.Println("/authorize")
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	log.Println(store.Get("ReturnUri"))

	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	r.Form = form
	log.Println(form)

	store.Delete("ReturnUri")
	store.Save()

	err = HandleAuthorizeRequest(w, r)
	if err != nil {
		log.Println(err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Println(w)
}
