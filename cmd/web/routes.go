package main

import (
	"net/http"
	"chatapp.new.net/ui"
)

func (app *application) routes() http.Handler{
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /rooms", app.roomList)
	mux.HandleFunc("GET /{room}", app.room)
	mux.HandleFunc("GET /messages/{room}", app.getMessages)
	mux.HandleFunc("POST /messages/{room}", app.postMessage)
	mux.HandleFunc("GET /user/signup",app.userSignup)
	mux.HandleFunc("POST /user/signup",app.userSignupPost)
	mux.HandleFunc("GET /user/login",app.userLogin)
	mux.HandleFunc("POST /user/login",app.userLoginPost)
	mux.HandleFunc("GET /user/logout",app.userLogoutPost)
	return app.sessionManager.LoadAndSave(mux)
}