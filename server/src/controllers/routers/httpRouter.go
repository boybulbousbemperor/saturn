package routers

import (
	"encoding/json"
	"net/http"

	"github.com/boybulbousbemperor/go-saturn/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

var params = []models.APIParameters{
	{Controller: "controller", Action: "action", ID: "id"},
}

func HttpServerNew(addr string) *http.Server {
	hand := mux.NewRouter()
	hand.HandleFunc("/", HandleProduce).Methods("POST")
	hand.HandleFunc("/", HandleConsume).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: hand,
	}
}

func HandleProduce(w http.ResponseWriter, r *http.Request) {
	ctx := *gin.Default()
	RouteParams(&gin.Context{
		Request:  r,
		Writer:   nil,
		Params:   []gin.Param{},
		Keys:     map[string]interface{}{},
		Errors:   []*gin.Error{},
		Accepted: []string{},
	})
	err := json.NewDecoder(r.Body).Decode(&ctx.HTMLRender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func HandleConsume(w http.ResponseWriter, r *http.Request) {
	ctx := *gin.Default()
	RouteParams(&gin.Context{
		Request:  r,
		Writer:   nil,
		Params:   []gin.Param{},
		Keys:     map[string]interface{}{},
		Errors:   []*gin.Error{},
		Accepted: []string{},
	})
	err := json.NewDecoder(r.Body).Decode(&ctx.HTMLRender)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func RouteParams(ctx *gin.Context) {
	cont := ctx.Param("controller")
	act := ctx.Param("action")
	id := ctx.Param("id")

	for _, x := range params {
		if x.Controller == cont {
			ctx.IndentedJSON(http.StatusOK, x)
		}

		if x.Action == act {
			ctx.IndentedJSON(http.StatusOK, x)
		}

		if x.ID == id {
			ctx.IndentedJSON(http.StatusOK, x)
			return
		}
	}
}
