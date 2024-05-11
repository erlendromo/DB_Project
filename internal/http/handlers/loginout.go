package handlers

import (
	"DB_Project/internal/http/middlewares"
	"DB_Project/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// LoginRequest struct
//
//	@title			LoginRequest
//	@description	This struct will be used to decode the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginLogoutResponse struct
//
//	@title			LoginLogoutResponse
//	@description	This struct will be used to encode the login/logout response
type LoginLogoutResponse struct {
	Message string `json:"message"`
}

// Login Log in
//
//	@title			Login
//	@summary		Log in
//	@description	Log in and set session
//	@tags			Login
//	@accept			json
//	@produce		json
//	@param			body	body		LoginRequest	true	"Login request"
//	@success		200		{object}	LoginLogoutResponse
//	@failure		500		{object}	utils.ErrorResponse
//	@router			/electromart/v1/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	if err := middlewares.SetSession(w, loginRequest.Username); err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	utils.JSON(w, http.StatusOK, LoginLogoutResponse{
		Message: fmt.Sprintf("You are logged in as %s!", loginRequest.Username),
	})

}

// Logout Log out
//
//	@title			Logout
//	@summary		Log out
//	@description	Log out and clear session (requires login)
//	@tags			Logout
//	@produce		json
//	@success		200	{object}	LoginLogoutResponse
//	@failure		401	{object}	utils.ErrorResponse
//	@router			/electromart/v1/logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
	statuscode, err := middlewares.ClearSession(w, r)
	if err != nil {
		utils.ERROR(w, statuscode, err)
		return
	}

	utils.JSON(w, http.StatusOK, LoginLogoutResponse{
		Message: "You are logged out!",
	})
}
