package middlewares

import (
	"net/http"

	"github.com/the-jey/gomushroomapi/errors"
	"github.com/the-jey/gomushroomapi/services"
	"github.com/the-jey/gomushroomapi/utils"
)

func IsLogin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tString := r.Header.Get("Authorization")
		if tString == "" {
			errors.SendJSONErrorResponse(w, "Please give a JWT token in the headers ❌", http.StatusUnauthorized)
			return
		}
		tString = tString[len("Bearer "):]

		if err := utils.VerifyJWTToken(tString); err != nil {
			errors.SendJSONErrorResponse(w, "Please give a valid token in the headers ❌", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tString := r.Header.Get("Authorization")
		if tString == "" {
			errors.SendJSONErrorResponse(w, "Please give a JWT token in the headers ❌", http.StatusUnauthorized)
			return
		}
		tString = tString[len("Bearer "):]

		if err := utils.VerifyJWTToken(tString); err != nil {
			errors.SendJSONErrorResponse(w, "Please give a valid token in the headers ❌", http.StatusUnauthorized)
			return
		}

		id, err := utils.ParseClaimsToken(tString)
		if err != nil {
			errors.SendJSONErrorResponse(w, "Error parsing the claims of the token ❌", http.StatusUnauthorized)
			return
		}

		u, s := services.GetUserByID(id)
		if s != "" {
			errors.SendJSONErrorResponse(w, s, http.StatusUnauthorized)
			return

		}

		if !u.IsAdmin {
			errors.SendJSONErrorResponse(w, "Need superior access rights ❌", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
