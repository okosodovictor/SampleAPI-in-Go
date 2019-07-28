package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func UsersRouter(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "users" {
		switch r.Method {
		case http.MethodGet:
			usersGetAll(w, r)
			return
		case http.MethodPost:
			usersPostOne(w, r)
			return
		case http.MethodHead:
			usersGetAll(w, r)
			return
		case http.MethodOptions:
			PostOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			PostError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "users/")
	if !bson.IsObjectIdHex(path) {
		PostError(w, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(path)
	switch r.Method {
	case http.MethodGet:
		usersGetOne(w, r, id)
		return
	case http.MethodPatch:
		usersPatchOne(w, r, id)
		return
	case http.MethodDelete:
		usersDeleteOne(w, r, id)
		return
	case http.MethodHead:
		usersGetOne(w, r, id)
		return
	case http.MethodOptions:
		PostOptionsResponse(w, []string{http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		PostError(w, http.StatusMethodNotAllowed)
		return
	}
}
