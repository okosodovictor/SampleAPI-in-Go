package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type jsonResponse map[string]interface{}

func PostError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func PostBodyResponse(w http.ResponseWriter, code int, content jsonResponse) {
	if content != nil {
		js, err := json.Marshal(content)

		if err != nil {
			PostError(w, http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(js)
		return
	}
	w.WriteHeader(code)
	w.Write([]byte(http.StatusText(code)))
}

func PostOptionsResponse(w http.ResponseWriter, methods []string, content jsonResponse) {
	w.Header().Set("Allow", strings.Join(methods, ","))
	PostBodyResponse(w, http.StatusOK, content)
}
