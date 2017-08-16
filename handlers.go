package main

import (
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
	"strconv"
)

func Root(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "Aloha!\n")
}

func GetUsers(w http.ResponseWriter, r *http.Request, params url.Values) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, params url.Values) {
	var id int
	var err error

	// Do type casting
	if id, err = strconv.Atoi(params["id"][0]); err != nil {
		panic(err)
	}

	user := RepoGetUser(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if user.Id > 0 {
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request, params url.Values) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	u := RepoCreateUser(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request, params url.Values) {
	var id int
	var err error

	// Do type casting
	if id, err = strconv.Atoi(params["id"][0]); err != nil {
		panic(err)
	}

	u := RepoGetUser(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if u.Id > 0 {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			panic(err)
		}

		defer r.Body.Close()

		user.Id = id
		u := RepoUpdateUser(user)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(u); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request, params url.Values) {
	var id int
	var err error

	// Do type casting
	if id, err = strconv.Atoi(params["id"][0]); err != nil {
		panic(err)
	}

	user := RepoGetUser(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if user.Id > 0 {
		RepoDeleteUser(id)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
