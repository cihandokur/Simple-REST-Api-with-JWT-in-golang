package model

import (
	"fmt"

	"github.com/cihandokur/simple_rest_api/helper"
)

type SingUp struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

func (s SingUp) Validate() error {

	if !helper.EmailCheck(s.Email) {
		return fmt.Errorf("invalid email")
	}
	if len(s.Password) < 4 {
		return fmt.Errorf("invalid password, Password should be more than 4 characters")
	}
	if len(s.FirstName) < 1 || len(s.LastName) < 1 {
		return fmt.Errorf("invalid firstname/lastname, please enter firstname/lastname")
	}

	return nil
}
