package main

import (
	"errors"
	"net/http"
	"strconv"
)

func (app *application) getQuizzesHandler(w http.ResponseWriter, r *http.Request) {
	language, err := strconv.ParseInt(r.URL.Query().Get("language"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid or missing language parameter"))
		return
	}

	quizzes, err := app.models.Quizzes.GetAll(language)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"quizzes": quizzes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
