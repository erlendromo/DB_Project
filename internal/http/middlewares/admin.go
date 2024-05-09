package middlewares

import (
	"DB_Project/internal/utils"
	"errors"
	"net/http"
)

func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessiondata, statuscode, err := GetUserFromSession(r)
		if err != nil {
			utils.ERROR(w, statuscode, utils.NewUnauthorizedError(err))
			return
		}

		if !sessiondata.Admin {
			utils.ERROR(w, http.StatusUnauthorized, utils.NewUnauthorizedError(errors.New("not an admin")))
			return
		}

		next(w, r)
	}
}
