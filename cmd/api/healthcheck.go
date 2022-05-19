package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	// comment this to return json string

	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment:%s\n", app.config.env)
	// fmt.Fprintf(w, "port:%d\n", app.config.port)

	// use a fixed format json string
	// js := `{"status":"available", "environment":%q, "version":%q}`
	// js = fmt.Sprintf(js, app.config.env, version)

	// use json encoder
	// create a map that we want to send
	data := envelope{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// pass the map to json marshall
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		// change into our error helper
		// app.logger.Println(err)
		// http.Error(w, "Server error and cannot processing your request", http.StatusInternalServerError)
		// return

		app.serverErrorResponse(w, r, err)
	}
}
