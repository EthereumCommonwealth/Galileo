package controllers

import "github.com/eduartua/callisto/Galileo/views"

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("mainLayout", "views/static/home.gohtml"),
		//Contact: views.NewView("mainLayout", "views/static/contac.gohtml"),
	}
}

type Static struct {
	Home *views.View
}