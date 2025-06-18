package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/tehweifu/languages4all-golang/internal/data"
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

func (app *application) savePointsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement cleaner function and validations

	var input struct {
		QuizID               int64 `json:"quiz_id"`
		UserID               int64 `json:"user_id"`
		Points               int   `json:"points"`
		Completed            int   `json:"completed"`
		CurrentQuestionOrder int   `json:"currentQuestionOrder"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	score := &data.Point{
		QuizID:               input.QuizID,
		UserID:               input.UserID,
		Points:               input.Points,
		Completed:            input.Completed,
		CurrentQuestionOrder: input.CurrentQuestionOrder,
	}

	err = app.models.Points.Upsert(score)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"saved": "ok"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
