package services

import (
	"net/http"
	"strings"

	"github.com/boybulbousbemperor/go-saturn/src/factory"
	"github.com/gin-gonic/gin"
)

type ContentService struct {
	Repository factory.ContentRepository
}

func (cs *ContentService) EncodingServiceGPPD(w http.ResponseWriter, r *http.Request) {
	content := strings.TrimPrefix(r.URL.Path, "api/")

	switch r.Method {
	case http.MethodGet:
		cs.GetContent(w, content)
	case http.MethodPost:
		cs.PostContent(w, content)
	}
}

func (cs *ContentService) GetContent(w http.ResponseWriter, content string) {
	cs.Repository.SetContent(content)
	if strings.Compare(content, w.Header().Get(content)) != 1 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusFound)
	}
}

func GetContentByID(id int) gin.Context {
	var context gin.Context
	return context
}

func (cs *ContentService) PostTitle(string, content string) {
	cs.Repository.SetContent(content)
}

func (cs *ContentService) PostContent(w http.ResponseWriter, content string) {
	cs.Repository.SetContent(content)
	w.WriteHeader(http.StatusAccepted)
}
