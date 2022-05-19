package main

import (
	"fmt"
	"net/http"
)

//generic helper to print error
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// create an errorResponse to return jsonerror
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// create a Server Error Response
// will be used when our app encounters error at runtime
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "The server encountered a problem and couldnt process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// create a not found error
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the request resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// create method not allowed response
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
