package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"

// 	"github.com/asdine/storm"
// 	"github.com/labstack/echo"
// 	"gopkg.in/mgo.v2/bson"
// )

// func usersOptions(c echo.Context) error {
// 	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
// 	c.Response().Header().Set("Allow", strings.Join(methods, ","))
// 	return c.NoContent(http.StatusOK)
// }

// func userOptions(c echo.Context) error {
// 	methods := []string{http.MethodGet, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
// 	c.Response().Header().Set("Allow", strings.Join(methods, ","))
// 	return c.NoContent(http.StatusOK)
// }

// func bodyToUser(r *http.Request, u *user.User) error {
// 	if r.Body == nil {
// 		return errors.New("Request body is empty")
// 	}

// 	if u == nil {
// 		return errors.New("A User is required")
// 	}

// 	bd, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	var result = json.Unmarshal(bd, u)
// 	fmt.Println(result)
// 	return result
// }

// func usersGetAll(w echo.Context, r *http.Request) {
// 	users, err := user.All()

// 	if err != nil {
// 		PostError(w, http.StatusInternalServerError)
// 		return
// 	}

// 	if r.Method == http.MethodHead {
// 		PostBodyResponse(w, http.StatusOK, jsonResponse{})
// 		return
// 	}

// 	PostBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
// }

// func usersPostOne(w echo.Context, r *http.Request) {

// 	u := new(user.User)
// 	fmt.Println(u)
// 	err := bodyToUser(r, u)
// 	if err != nil {
// 		PostError(w, http.StatusBadRequest)
// 		return
// 	}

// 	u.ID = bson.NewObjectId()
// 	fmt.Println(u.ID)
// 	err = u.Save()

// 	if err != nil {
// 		if err == user.ErrRecordInvalid {
// 			PostError(w, http.StatusBadRequest)
// 		} else {
// 			PostError(w, http.StatusInternalServerError)
// 		}

// 		return
// 	}

// 	w.Header().Set("Location", "/users/"+u.ID.Hex())
// 	w.WriteHeader(http.StatusCreated)
// }

// func usersGetOne(c echo.Context) {
// 	u, err := user.One(id)
// 	if err != nil {
// 		if err == storm.ErrNotFound {
// 			PostError(w, http.StatusNotFound)
// 			return
// 		}
// 		PostError(w, http.StatusInternalServerError)
// 		return
// 	}

// 	if c.Request().Method == http.MethodHead {
// 		PostBodyResponse(w, http.StatusOK, jsonResponse{})
// 		return
// 	}

// 	PostBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
// }
// func usersPutOne(w echo.Context, r *http.Request) {

// 	u := new(user.User)
// 	fmt.Println(u)
// 	err := bodyToUser(r, u)
// 	if err != nil {
// 		PostError(w, http.StatusBadRequest)
// 		return
// 	}

// 	u.ID = id
// 	err = u.Save()

// 	if err != nil {
// 		if err == user.ErrRecordInvalid {
// 			PostError(w, http.StatusBadRequest)
// 		} else {
// 			PostError(w, http.StatusInternalServerError)
// 		}

// 		return
// 	}

// 	PostBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
// }

// func usersPatchOne(c echo.Context, r *http.Request) {

// 	u, err := user.One(id)
// 	if err != nil {
// 		if err == storm.ErrNotFound {
// 			return echo.NewHTTPError(http.StatusNotFound)
// 		}
// 		return echo.NewHTTPError(http.StatusInternalServerError)
// 	}

// 	err = bodyToUser(r, u)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest)
// 	}

// 	u.ID = id
// 	err = u.Save()

// 	if err != nil {
// 		if err == user.ErrRecordInvalid {
// 			return echo.NewHTTPError(http.StatusBadRequest)
// 		} else {
// 			return echo.NewHTTPError(http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	PostBodyResponse(c, http.StatusOK, jsonResponse{"user": u})
// }

// func usersDeleteOne(c echo.Context) {
// 	err := user.Delete(id)
// 	if err != nil {
// 		if err == storm.ErrNotFound {
// 			return echo.NewHTTPError(c, http.StatusNotFound)
// 		}
// 		return echo.NewHTTPError(http.StatusInternalServerError)
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Running API V1")
}

func main() {
	e := echo.New()
	e.GET("/", root)
}
// 	u := e.Group("/users")
// 	u.OPTIONS("", usersOptions)
// 	u.HEAD("", usersGetAll)
// 	u.GET("", usersGetAll)
// 	u.GET("", usersPostOne)

// 	uid := u.Group("/:id")

// 	uid.OPTIONS("", userOptions)
// 	uid.HEAD("", usersGetOne)
// 	uid.GET("", usersGetOne)
// 	uid.Put("", usersPutOne)
// 	uid.PATCH("", usersPatchOne)
// 	uid.DELETE("", usersDeleteOne)

// 	e.Start(":1323")
// }
