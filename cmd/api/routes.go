package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/languages", app.getLanguagesHandler)

	router.HandlerFunc(http.MethodGet, "/v1/points", app.getPointsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/points", app.savePointsHandler)

	router.HandlerFunc(http.MethodGet, "/v1/questions", app.getQuestionsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/questions/random", app.getRandomQuestionsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/questions/count", app.getQuestionCountHandler)

	router.HandlerFunc(http.MethodGet, "/v1/quizzes", app.getQuizzesHandler)

	router.HandlerFunc(http.MethodGet, "/v1/users", app.getUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/ranking", app.getUserRankingHandler)

	router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	return app.metrics(app.recoverPanic(app.rateLimit(router)))
}
