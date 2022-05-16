package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	// when httprouter is parsing a request, any interpolated URL params will be
	// stored in the request context. We can use the paramsFromContext() function
	// to retrieve a slice containing

	params := httprouter.ParamsFromContext(r.Context())

	// we can use ByName() method to get the value "id" parameter from the slice
	// but the value of id will always returning string
	// so we try to convert it into integer
	// if the parameter is invalid and couldnt be converted, id is invalid
	// if id is invalid return http.NotFound()

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid parameter")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// js, err := json.Marshal(data)

	// use MarshalIndent to add indent, looks cleaner in console and browser
	// NOTE: MarshalIndent() took 65% longer and 30% more memory than Marshal()
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	// set content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
