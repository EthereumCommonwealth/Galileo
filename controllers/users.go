package controllers

import (
	"net/http"
	"fmt"

	"github.com/eduartua/callisto/Galileo/views"
	"github.com/eduartua/callisto/Galileo/models"
	"github.com/eduartua/callisto/Galileo/rand"
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
	err := u.signIn(w, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Redirect to the cookie test page to test the cookie
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

// Login is used to process the login form
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address.")
		case models.ErrIvalidPassword:
			fmt.Fprintln(w, "Invalid password provided.")
		case nil:
			fmt.Fprintln(w, user)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	err = u.signIn(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

//Sign the given user in via cookies
func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
		if err != nil {
			return err
		}
	}
	cookie := http.Cookie{
		Name: "remember_token",
		Value: user.Remember,
	}
	http.SetCookie(w, &cookie)
	return nil
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