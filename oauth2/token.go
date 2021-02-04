package oauth2

import (
	"net/http"
)

func Token(w http.ResponseWriter, r *http.Request) {
	err := HandleTokenRequest(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
