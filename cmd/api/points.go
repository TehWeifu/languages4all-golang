package main

import (
	"errors"
	"net/http"
	"strconv"
)

func (app *application) getPointsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement cleaner function and validations
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid or missing user_id parameter"))
		return
	}
	language, err := strconv.ParseInt(r.URL.Query().Get("language"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid or missing language parameter"))
		return
	}

	points, err := app.models.Points.GetAll(userId, language)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"points": points}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
