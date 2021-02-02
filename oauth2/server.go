package oauth2

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-session/session"
)

var (
	gServer *server.Server
	once    sync.Once
)

func InitServer(manager *manage.Manager) *server.Server {
	once.Do(func() {
		gServer = server.NewServer(server.NewConfig(), manager)
		gServer.SetResponseErrorHandler(responseErrorHandler)
		gServer.SetInternalErrorHandler(internalErrorHandler)
		gServer.SetUserAuthorizationHandler(userAuthorizeHandler)
	})
	return gServer
}

func HandleAuthorizeRequest(w http.ResponseWriter, r *http.Request) error {
	return gServer.HandleAuthorizeRequest(w, r)
}

func HandleTokenRequest(w http.ResponseWriter, r *http.Request) error {
	return gServer.HandleTokenRequest(w, r)
}

func ValidationBearerToken(r *http.Request) (oauth2.TokenInfo, error) {
	return gServer.ValidationBearerToken(r)
}

func responseErrorHandler(re *errors.Response) {
	log.Println("Response Error:", re.Error.Error())
}

func internalErrorHandler(err error) (re *errors.Response) {
	log.Println("Internal Error:", err.Error())
	return
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	log.Println("userAuthorizeHandler")

	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	log.Println("userAuthorizeHandler", uid)

	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		log.Println("userAuthorizeHandler", r.Form)

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}
