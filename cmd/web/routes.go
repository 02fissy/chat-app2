package main

import (
	"net/http"
	"chatapp.new.net/ui"
)

func (app *application) routes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.HandleFunc("GET /", app.roomList)
	mux.HandleFunc("GET /{room}", app.home)
	mux.HandleFunc("GET /messages/{room}", app.getMessages)
	mux.HandleFunc("POST /messages/{room}/users/{userID}", app.postMessage)
	return mux
}