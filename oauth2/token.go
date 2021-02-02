package oauth2

import (
	"log"
	"net/http"
)

func Token(w http.ResponseWriter, r *http.Request) {
	log.Println("/token")

	err := HandleTokenRequest(w, r)
	log.Println(w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
