package controllers

import (
	"github.com/eduartua/callisto/Galileo/views"
	"net/http"
	"fmt"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("mainLayout", "views/users/new.gohtml"),
	}
}

func (u *Users ) New(w http.ResponseWriter, r *http.Request)  {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users ) Create(w http.ResponseWriter, r *http.Request)  {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is:", form.Email)
	fmt.Fprintln(w, "Password is:", form.Password)
}

type Users struct {
	NewView *views.View
}

//This is going to change soon
type SignupForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}