package main

import (
	"errors"
	"net/http"
	"strconv"
)

func (app *application) getQuestionCountHandler(w http.ResponseWriter, r *http.Request) {
	language, err := strconv.ParseInt(r.URL.Query().Get("language"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid or missing language parameter"))
		return
	}

	count, err := app.models.Questions.GetCountByLanguage(language)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"count_questions": count}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	quizID, err := strconv.ParseInt(r.URL.Query().Get("quiz_id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid or missing quiz_id parameter"))
		return
	}

	questions, err := app.models.Questions.GetAll(quizID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"questions": questions}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getRandomQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	languageID, err := strconv.ParseInt(r.URL.Query().Get("language"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid or missing language parameter"))
		return
	}

	questions, err := app.models.Questions.GetRandomByLanguage(languageID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"questions": questions}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
