package main

import (
	"net/http"
)

func (app *application) getLanguagesHandler(w http.ResponseWriter, r *http.Request) {
	languages, err := app.models.Languages.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"languages": languages}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
