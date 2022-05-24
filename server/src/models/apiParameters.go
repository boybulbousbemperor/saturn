package models

type APIParameters struct { // Controller, Action, ID parameter(s)
	Controller string `json:controller`
	Action     string `json:action`
	ID         string `json:id`
}
