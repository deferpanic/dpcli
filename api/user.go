package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct{}

type User struct {
	Email    string
	Username string
	Password string
}

type UsersResponse struct {
	Token string
	Error string
}

// Create creates a new user and returns the token
func (users *Users) Create(email string, username string, password string) {

	user := &User{}
	user.Email = email
	user.Username = username
	user.Password = password

	b, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	ur := UsersResponse{}
	err = Cli.PostJSON(b, APIBase+"/users/create", &ur)

	if err != nil {
		fmt.Println(RedBold(err.Error()))
		os.Exit(1)
	} else {
		if ur.Error != "" {
			fmt.Println(RedBold(ur.Error))
			os.Exit(1)
		} else {
			fmt.Println(GreenBold("Successfully created user."))
			err = ioutil.WriteFile(os.Getenv("HOME")+"/.dprc", []byte(ur.Token), 0644)
			if err != nil {
				fmt.Println(RedBold(err.Error()))
				os.Exit(1)
			}
		}

	}
}
