package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func getURL(r *http.Request) string {
	ctx := r.Context()
	return ctx.Value("URL").(string)
}

func httpError(w http.ResponseWriter, err error) bool {
	if err != nil {
		// print the error to the stdout
		log.Println(err.Error())

		// write the error to the response
		http.Error(w, err.Error(), 500)
		return true
	}
	return false
}

func writeJSON(w http.ResponseWriter, data interface{}) error {
	// encode the response data to json
	jData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// set the json response content-type header
	w.Header().Set("Content-Type", "application/json")
	// write the data to the response
	w.Write(jData)
	return nil
}
