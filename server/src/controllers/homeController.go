package controllers

import (
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method == "GET" {

		file, handler, err := r.FormFile(`json:action`) // or have a typedef struct with `json:action` as a member
		if file != nil {
			return 1, file.Close()
		}
		if handler != nil {
			return 1, nil
		}
		if err != nil {
			return 1, err
		}

		id := r.URL.Path[len("/home/index"):]

		wrt, err := w.Write([]byte(id))
		if err != nil {
			return 1, err
		}

		return wrt, err
	}

	return 0, nil
}

func GetPrivacy(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method == "GET" {

		file, handler, err := r.FormFile(`json:action`) // or have a typedef struct with `json:action` as a member
		if file != nil {
			return 1, file.Close()
		}
		if handler != nil {
			return 1, nil
		}
		if err != nil {
			return 1, err
		}

		id := r.URL.Path[len("/home/privacy"):]

		wrt, err := w.Write([]byte(id))
		if err != nil {
			return 1, err
		}

		return wrt, err
	}

	return 0, nil
}
