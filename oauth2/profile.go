package oauth2

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	log.Println("/profile")

	token, err := ValidationBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("/test", token)

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":  token.GetClientID(),
		"user_id":    token.GetUserID(),
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(data)
	log.Println("/test", token)

}
