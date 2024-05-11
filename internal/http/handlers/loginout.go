package handlers

import (
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/http/middlewares"
	"DB_Project/internal/utils"
	"encoding/json"
	"errors"
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
//	@title			LoginResponse
//	@description	This struct will be used to encode the login response
type LoginResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// LogoutResponse struct
//
//	@title			LogoutResponse
//	@description	This struct will be used to encode the logout response
type LogoutResponse struct {
	Message string `json:"message"`
}

// Login Log in
//
//	@title			Login
//	@summary		Login
//	@description	Login and set session
//	@tags			Login
//	@accept			json
//	@produce		json
//	@param			body	body		LoginRequest	true	"Login request"
//	@success		200		{object}	LoginResponse
//	@failure		400		{object}	utils.ErrorResponse
//	@failure		500		{object}	utils.ErrorResponse
//	@router			/electromart/v1/customers/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	customer, err := dependencies.Dependencies.CustomerAddressDeps.PSQLCustomer.GetCustomerByUsername(r.Context(), loginRequest.Username)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, errors.New("invalid username or password"))
		return
	}

	if customer.Password != loginRequest.Password {
		utils.ERROR(w, http.StatusBadRequest, errors.New("invalid username or password"))
		return
	}

	if err := middlewares.SetSession(w, loginRequest.Username); err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	var role string
	if customer.Role == 1 {
		role = "admin"
	} else {
		role = "default_customer"
	}

	utils.JSON(w, http.StatusOK, LoginResponse{
		UserID:   fmt.Sprint(customer.ID),
		Username: loginRequest.Username,
		Role:     role,
	})

}

// Logout Log out
//
//	@title			Logout
//	@summary		Log out
//	@description	Log out and clear session (requires login)
//	@tags			Logout
//	@produce		json
//	@success		200	{object}	LogoutResponse
//	@failure		401	{object}	utils.ErrorResponse
//	@router			/electromart/v1/customers/logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
	statuscode, err := middlewares.ClearSession(w, r)
	if err != nil {
		utils.ERROR(w, statuscode, err)
		return
	}

	utils.JSON(w, http.StatusOK, LogoutResponse{
		Message: "You are logged out",
	})
}
