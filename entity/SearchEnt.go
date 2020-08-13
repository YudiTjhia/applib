package entity

import "applib/app"

type SearchEnt struct {
	app.RequestHeader
	Field string `json:"field"`
	Value string `json:"value"`
}

type SearchEnt2 struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
