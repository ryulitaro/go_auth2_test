package main

import (
	"log"
	"net/http"

	"mod.lge.com/code/projects/CLOUD/repos/poc-top-oauth2/oauth2"
)

func main() {
	manager, _ := oauth2.NewManager()
	oauth2.InitServer(manager)

	http.HandleFunc("/login", oauth2.LoginHandler)
	http.HandleFunc("/auth", oauth2.AuthHandler)

	http.HandleFunc("/authorize", oauth2.Authorize)

	http.HandleFunc("/token", oauth2.Token)

	http.HandleFunc("/profile", oauth2.Profile)

	log.Println("Server is running at 9096 port.")
	log.Fatal(http.ListenAndServe(":9096", nil))
}
