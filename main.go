package main

import (
	"net/http"
	"github.com/eduartua/callisto/Galileo/blocks"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"github.com/eduartua/callisto/Galileo/users"
	"github.com/eduartua/callisto/Galileo/config"
)

//I haven't implemented any database yet for Users and Sessions
var dbUsers = map[string]users.User{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/blocks", blocks.Index)
	http.HandleFunc("/block/:blockId", blocks.BlockInfo)
	http.HandleFunc("/tx/:transactionId", blocks.TxInfo)
	http.HandleFunc("/address/:addressId", blocks.AddressId)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/blocks", http.StatusSeeOther)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u users.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	config.TPL.ExecuteTemplate(w, "login.gohtml", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}