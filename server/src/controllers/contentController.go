package controllers

import (
	"html/template"
	"net/http"

	services "github.com/boybulbousbemperor/go-saturn/src/models/services"
	"github.com/gin-gonic/gin"
)

func CreateContent(w http.ResponseWriter, r *http.Request) (int, error) {
	cs := services.ContentService{}

	if r.Method == "POST" {
		r.ParseForm()

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

		r.Form.Add("/create", "/id")

		id := r.URL.Path[len("/create/"):]
		title := template.HTMLEscapeString(r.Form.Get("title"))
		content := template.HTMLEscapeString(r.Form.Get(string(id)))
		cs.PostTitle(title, content)

		wrt, err := w.Write([]byte(id))
		if err != nil {
			return 1, err
		}

		return wrt, err
	}

	return 0, nil
}

func EditContent(g *gin.Context, r *http.Request) {
	index := r.RequestURI
	services.GetContentByID(g.GetInt(index))

	if r.Method == "GET" {
		r.ParseForm()
		id := r.URL.Path[len("/edit/"):]

		file, handler, err := r.FormFile(`json:action`) // or have a typedef struct with `json:action` as a member
		if file != nil {
			file.Close()
		}
		if handler != nil {
			handler.Header.Add(id, g.HandlerName()) // make api for handling missing or non-existent files
			print(err.Error())
		}
		if err != nil {
			print(err.Error())
		}

		r.Form.Set("/edit", "/id")
	}
}
