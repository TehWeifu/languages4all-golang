package main

import (
	"errors"
	"net/http"

	"github.com/tehweifu/languages4all-golang/internal/data"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	androidUid := r.URL.Query().Get("android_id")
	if androidUid == "" {
		app.badRequestResponse(w, r, errors.New("invalid or missing android_id parameter"))
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		app.badRequestResponse(w, r, errors.New("invalid or missing name parameter"))
		return
	}

	user, err := app.models.Users.GetByNameAndAndroidUid(name, androidUid)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
