package oauth2

import (
	"encoding/json"
	"net/http"
	"time"
)

func TokenInfo(w http.ResponseWriter, r *http.Request) {
	token, err := ValidationBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"scope":      token.GetScope(),
		"userId":     token.GetUserID(),
		"clientId":   token.GetClientID(),
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(data)
}
