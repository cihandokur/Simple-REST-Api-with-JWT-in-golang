package middleware

import (
	"net/http"

	"github.com/cihandokur/devlab/config"
	"github.com/cihandokur/devlab/helper"
)

var apiError = ApiError{}

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		headerToken := r.Header.Get(config.Config.Jwt.Header)

		if headerToken == "" {
			apiError.HandleErr(w, http.StatusForbidden, "Token not provided!")
			return
		}

		_, err := helper.ValidateToken(headerToken)
		if err != nil {
			apiError.HandleErr(w, http.StatusForbidden, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
