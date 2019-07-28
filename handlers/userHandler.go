package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asdine/storm"

	user "../users"
	"gopkg.in/mgo.v2/bson"
)

func bodyToUser(r *http.Request, u *user.User) error {
	if r.Body == nil {
		return errors.New("Request body is empty")
	}

	if u == nil {
		return errors.New("A User is required")
	}

	bd, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		return err
	}
	var result = json.Unmarshal(bd, u)
	fmt.Println(result)
	return result
}

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()

	if err != nil {
		PostError(w, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodHead {
		PostBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}

	PostBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}

func usersPostOne(w http.ResponseWriter, r *http.Request) {

	u := new(user.User)
	fmt.Println(u)
	err := bodyToUser(r, u)
	if err != nil {
		PostError(w, http.StatusBadRequest)
		return
	}

	u.ID = bson.NewObjectId()
	fmt.Println(u.ID)
	err = u.Save()

	if err != nil {
		if err == user.ErrRecordInvalid {
			PostError(w, http.StatusBadRequest)
		} else {
			PostError(w, http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Location", "/users/"+u.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}

func usersGetOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			PostError(w, http.StatusNotFound)
			return
		}
		PostError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		PostBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}

	PostBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}
func usersPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {

	u := new(user.User)
	fmt.Println(u)
	err := bodyToUser(r, u)
	if err != nil {
		PostError(w, http.StatusBadRequest)
		return
	}

	u.ID = id
	err = u.Save()

	if err != nil {
		if err == user.ErrRecordInvalid {
			PostError(w, http.StatusBadRequest)
		} else {
			PostError(w, http.StatusInternalServerError)
		}

		return
	}

	PostBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}

func usersPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {

	u, err := user.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			PostError(w, http.StatusNotFound)
			return
		}
		PostError(w, http.StatusInternalServerError)
		return
	}

	err = bodyToUser(r, u)
	if err != nil {
		PostError(w, http.StatusBadRequest)
		return
	}

	u.ID = id
	err = u.Save()

	if err != nil {
		if err == user.ErrRecordInvalid {
			PostError(w, http.StatusBadRequest)
		} else {
			PostError(w, http.StatusInternalServerError)
		}

		return
	}

	PostBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}

func usersDeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := user.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			PostError(w, http.StatusNotFound)
			return
		}
		PostError(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
