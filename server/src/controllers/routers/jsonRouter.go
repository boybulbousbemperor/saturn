package routers

import (
	"encoding/json"
	"net/http"

	"github.com/boybulbousbemperor/go-saturn/src/models"
)

func routeJSON() {
	mux := http.NewServeMux()
	mux.HandleFunc("api/", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp := models.APIParameters{
			Controller: "controller",
			Action:     "action",
			ID:         "id",
		}
		if err := enc.Encode(resp); err != nil {
			panic("Encoding of reponse: failed!")
		}
		HandleProduce(w, r)
		HandleConsume(w, r)
	})
}
