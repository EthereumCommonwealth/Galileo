package main

import (
	"fmt"
	"github.com/eduartua/callisto/Galileo/models"
)

const (
host     = "localhost"
port     = 5432
user     = "postgres"
password = "hola"
dbname   = "galileo_dev"
)

func main() {
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
"password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
us, err := models.NewUserService(psqlInfo)
if err != nil {
panic(err)
}
defer us.Close()
us.DestructiveReset()

user := models.User{
Name:     "Michael Scott",
Email:    "michael@dundermifflin.com",
Password: "bestboss",
}
err = us.Create(&user)
if err != nil {
panic(err)
}
// Verify that the user has a Remember and RememberHash
fmt.Printf("%+v\n", user)
if user.Remember == "" {
panic("Invalid remember token")
}

// Now verify that we can lookup a user with that remember
// token
user2, err := us.ByRemember(user.Remember)
if err != nil {
panic(err)
}
fmt.Printf("%+v\n", *user2)
}