package handlers

import "net/http"

func Default(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"Viesti": "Moi vaan!"}

	// get datastore handler
	//db := getStore(r)

	// print response to the http writer
	err := writeJSON(w, data)
	httpError(w, err)
}
