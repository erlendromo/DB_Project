package handlers

import (
	"DB_Project/internal/http/middlewares"
	"DB_Project/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginLogoutResponse struct {
	Message string `json:"message"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	middlewares.SetSession(w, 6969, loginRequest.Username)

	utils.JSON(w, http.StatusOK, LoginLogoutResponse{
		Message: fmt.Sprintf("You are logged in as %s!", loginRequest.Username),
	})

}

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

func MyProfile(w http.ResponseWriter, r *http.Request) {
	sessiondata, statuscode, err := middlewares.GetUserFromSession(r)
	if err != nil {
		utils.ERROR(w, statuscode, utils.NewUnauthorizedError(err))
		return
	}

	// TODO get customer from database
	_ = sessiondata

}

func UpdateMyProfile(w http.ResponseWriter, r *http.Request) {
	sessiondata, statuscode, err := middlewares.GetUserFromSession(r)
	if err != nil {
		utils.ERROR(w, statuscode, utils.NewUnauthorizedError(err))
		return
	}

	// TODO get customer from database
	_ = sessiondata
}

func DeleteMyProfile(w http.ResponseWriter, r *http.Request) {
	sessiondata, statuscode, err := middlewares.GetUserFromSession(r)
	if err != nil {
		utils.ERROR(w, statuscode, utils.NewUnauthorizedError(err))
		return
	}

	// TODO get customer from database
	_ = sessiondata
}

// AllCustomers Get all customers
//
//	@summary		Get all customers
//	@description	Get all customers
//	@tags			Customer
//	@security		AdminAuth
//	@produce		json
//	@success		200	{json}		message
//	@failure		401	{object}	utils.ErrorResponse
//	@router			/electromart/v1/customers [get]
//	@router			/electromart/v1/customers/ [get]
func AllCustomers(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Message string `json:"message"`
	}{
		Message: "You are authorized!",
	}

	utils.JSON(w, http.StatusOK, resp)
}
