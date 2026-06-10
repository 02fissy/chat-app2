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
func (app *application) getMessages(
	w http.ResponseWriter,
	r *http.Request,
) {

	roomName := r.PathValue("room")

	roomID, err := app.rooms.GetOrCreate(roomName)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	messages, err :=
		app.messages.GetByRoomID(roomID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set(
		"Content-Type",
		"text/plain",
	)

	for _, msg := range messages {

		fmt.Fprintf(
			w,
			"%d|%s\n",
			msg.UserID,
			msg.Content,
		)
	}
}
func (app *application) postMessage(
	w http.ResponseWriter,
	r *http.Request,
) {

	roomName := r.PathValue("room")
	
	roomID, err :=
		app.rooms.GetOrCreate(roomName)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	userID, err :=
		strconv.Atoi(
			r.PathValue("userID"),
		)

	if err != nil {
		http.Error(w, "invalid user", 400)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	content := r.FormValue("message")

	err = app.messages.Insert(
		roomID,
		userID,
		content,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}