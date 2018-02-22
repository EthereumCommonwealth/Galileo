package controllers

import (
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/eduartua/callisto/Galileo/models"
	"github.com/eduartua/callisto/Galileo/rand"
)

func NewUsers(us *models.UserService) *Users {
	return &Users{
		us: us,
	}
}


//This code is in dev mode
func (u *Users) Create(w http.ResponseWriter, r *http.Request) () {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	var user models.User
	json.Unmarshal(body, &user)
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := u.signIn(w, &user)
	if err != nil {
		// Temporarily render the error message for debugging
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	return

}

func (u *Users ) Login(w http.ResponseWriter, r *http.Request) {
	dataForm := DataForm{}
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	json.Unmarshal(body, &dataForm)

	user, err := u.us.Authenticate(dataForm.Email, dataForm.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address.")
		case models.ErrInvalidPassword:
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
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}

func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := u.us.ByRemember(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)
}

type Users struct {
	us 			*models.UserService
}

//This is going to change soon
type SignupForm struct {
	Name string	`schema: "name"`
	Email string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email		string `schema:"email"`
	Password	string `schema:"password"`
}

type DataForm struct {
	Email		string `json:"email"`
	Password	string `json:"password"`
}