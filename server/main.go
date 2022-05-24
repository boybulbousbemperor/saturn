package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type HubComponent struct {
	app.Compo
}

func (hubCompo HubComponent) RenderView() app.UI {
	return app.H1().Text("Welcome to Saturn!")
}

func main() {
	app.Route("/", &app.Compo{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "$Name",
		Description: "$Description",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
