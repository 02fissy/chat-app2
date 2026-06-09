package main

import (
	"fmt"
	"net/http"
	"strconv"
)
func (app *application) home(w http.ResponseWriter, r *http.Request){
	room := r.PathValue("room")

	data := struct {
		Room string
	}{
		Room: room,
	}

	app.render(w, r, http.StatusOK, "room.html", data)
}
func (app *application) getMessages(w http.ResponseWriter, r *http.Request) {

	room := r.PathValue("room")

	roomID, err := app.rooms.GetByName(room)
	if err != nil {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	messages, err := app.messages.GetByRoomID(roomID)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	for _, m := range messages {
		fmt.Fprintf(w, "%d|%s\n", m.UserID, m.Content)
	}
}
func (app *application) postMessage(w http.ResponseWriter, r *http.Request) {

	room := r.PathValue("room")
	userID, err := strconv.Atoi(r.PathValue("userID"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	roomID, err := app.rooms.GetByName(room)
	if err != nil {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form", http.StatusBadRequest)
		return
	}

	content := r.FormValue("message")
	if content == "" {
		http.Error(w, "empty message", http.StatusBadRequest)
		return
	}

	err = app.messages.Insert(roomID, userID, content)
	if err != nil {
		http.Error(w, "failed to insert", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ok"))
}