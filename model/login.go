package model

import (
	"fmt"

	"github.com/cihandokur/simple_rest_api/helper"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l Login) Validate() error {

	if !helper.EmailCheck(l.Email) {
		return fmt.Errorf("invalid email")
	}
	if len(l.Password) < 4 {
		return fmt.Errorf("invalid password, Password should be more than 4 characters")
	}

	return nil
}
