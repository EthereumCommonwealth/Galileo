package controllers

import (
	"net/http"
	"fmt"

	"github.com/eduartua/callisto/Galileo/views"
	"github.com/eduartua/callisto/Galileo/models"
)

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("mainLayout", "views/users/new.gohtml"),
		us: us,
	}
}

func (u *Users ) New(w http.ResponseWriter, r *http.Request)  {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//This code is in dev mode
func (u *Users ) Create(w http.ResponseWriter, r *http.Request)  {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name: form.Name,
		Email: form.Email,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "User is", user)
}

type Users struct {
	NewView *views.View
	us 		*models.UserService
}

//This is going to change soon
type SignupForm struct {
	Name string	`schema: "name"`
	Email string `schema:"email"`
	Password string `schema:"password"`
}