package main

import (
	"net/http"
)

func (app *application) listLanguagesHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *application) getLanguageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		// Use the new notFoundResponse() helper.
		app.notFoundResponse(w, r)
		return
	}

	language, err := app.models.Languages.GetById(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"language": language}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
