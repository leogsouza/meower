package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/leogsouza/meower/db"
	"github.com/leogsouza/meower/event"
	"github.com/leogsouza/meower/schema"

	"github.com/leogsouza/meower/util"
	"github.com/segmentio/ksuid"
)

func createMeowHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID string `json:"id"`
	}

	ctx := r.Context()

	// Read parameters
	body := template.HTMLEscapeString(r.FormValue("body"))
	log.Println("Creating new meow")
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	// Create meow
	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandomWithTime(createdAt)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create meow")
		return
	}

	meow := schema.Meow{
		ID:        id.String(),
		Body:      body,
		CreatedAt: createdAt,
	}

	if err := db.InsertMeow(ctx, meow); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create meow")
		return
	}

	// Publish event
	if err := event.PublishMeowCreated(meow); err != nil {
		log.Println(err)
	}

	// Return new meow
	util.ResponseOk(w, response{ID: meow.ID})
}
