package handlers

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/http/middlewares"
	"DB_Project/internal/utils"
	"encoding/json"
	"net/http"
)

// Signup Create a new customer
//
//	@title			Signup
//	@summary		Create a new customer
//	@description	Create a new customer
//	@tags			Customer
//	@accept			json
//	@produce		json
//	@param			body	body		customeraddressdomain.CreateCustomerAddressRequest	true	"Create customer"
//	@success		201		{object}	customeraddressdomain.DBCustomerAddress
//	@failure		422		{object}	utils.ErrorResponse
//	@failure		500		{object}	utils.ErrorResponse
//	@router			/electromart/v1/customers/signup [post]
func Signup(w http.ResponseWriter, r *http.Request) {
	var signuprequest customeraddressdomain.CreateCustomerAddressRequest
	if err := json.NewDecoder(r.Body).Decode(&signuprequest); err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	if errs := signuprequest.Validate(); len(errs) > 0 {
		utils.JSON(w, http.StatusUnprocessableEntity, utils.NewValidateErrors(errs))
		return
	}

	customer, err := dependencies.Dependencies.CustomerAddressDeps.PSQLCustomerAddress.CreateCustomerAddress(r.Context(), &signuprequest)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	utils.JSON(w, http.StatusCreated, customer)
}

// MyProfile Get my profile
//
//	@title			MyProfile
//	@summary		Get my profile
//	@description	Get my profile (requires login)
//	@tags			Customer
//	@security		UserAuth
//	@produce		json
//	@success		200	{object}	customeraddressdomain.CustomerAddresses
//	@failure		401	{object}	utils.ErrorResponse
//	@failure		500	{object}	utils.ErrorResponse
//	@router			/electromart/v1/customers/me [get]
func MyProfile(w http.ResponseWriter, r *http.Request) {
	sessiondata, statuscode, err := middlewares.GetUserFromSession(r)
	if err != nil {
		utils.ERROR(w, statuscode, utils.NewUnauthorizedError(err))
		return
	}

	cd := dependencies.Dependencies.CustomerAddressDeps.PSQLCustomerAddress

	customer, err := cd.GetCustomerAddressesByCustomerID(r.Context(), sessiondata.ID)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	utils.JSON(w, http.StatusOK, customer)
}

// UpdateMyProfile Update my profile
//
//	@title			UpdateMyProfile
//	@summary		Update my profile
//	@description	Update my profile (requires login)
//	@tags			Customer
//	@security		UserAuth
//	@accept			json
//	@produce		json
//	@param			body	body		customeraddressdomain.CreateCustomerAddressRequest	true	"Update customer"
//	@success		200		{object}	customeraddressdomain.CreateCustomerAddressRequest
//	@failure		401		{object}	utils.ErrorResponse
//	@failure		422		{object}	utils.ErrorResponse
//	@failure		500		{object}	utils.ErrorResponse
//	@router			/electromart/v1/customers/me [put]
func UpdateMyProfile(w http.ResponseWriter, r *http.Request) {
	sessiondata, statuscode, err := middlewares.GetUserFromSession(r)
	if err != nil {
		utils.ERROR(w, statuscode, utils.NewUnauthorizedError(err))
		return
	}

	var updatecustomer customeraddressdomain.CreateCustomerAddressRequest
	if err := json.NewDecoder(r.Body).Decode(&updatecustomer); err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	if errs := updatecustomer.Validate(); len(errs) > 0 {
		utils.JSON(w, http.StatusUnprocessableEntity, utils.NewValidateErrors(errs))
		return
	}

	// TODO fix this (is scuffed atm)

	if err = dependencies.Dependencies.CustomerAddressDeps.PSQLCustomer.UpdateCustomer(r.Context(), sessiondata.ID, &customeraddressdomain.CreateCustomer{
		Username:    updatecustomer.Username,
		Password:    updatecustomer.Password,
		FirstName:   updatecustomer.FirstName,
		LastName:    updatecustomer.LastName,
		Email:       updatecustomer.Email,
		PhoneNumber: updatecustomer.PhoneNumber,
	}); err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	utils.JSON(w, http.StatusOK, updatecustomer)
}

// DeleteMyProfile Delete my profile
//
//	@title			DeleteMyProfile
//	@summary		Delete my profile
//	@description	Delete my profile (requires login)
//	@tags			Customer
//	@security		UserAuth
//	@produce		json
//	@success		204
//	@failure		401	{object}	utils.ErrorResponse
//	@failure		500	{object}	utils.ErrorResponse
//	@router			/electromart/v1/customers/me [delete]
func DeleteMyProfile(w http.ResponseWriter, r *http.Request) {
	sessiondata, statuscode, err := middlewares.GetUserFromSession(r)
	if err != nil {
		utils.ERROR(w, statuscode, utils.NewUnauthorizedError(err))
		return
	}

	if err = dependencies.Dependencies.CustomerAddressDeps.PSQLCustomer.SoftDeleteCustomer(r.Context(), sessiondata.ID); err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AllCustomers Get all customers
//
//	@title			AllCustomers
//	@summary		Get all customers
//	@description	Get all customers (requires admin login)
//	@tags			Customer
//	@security		AdminAuth
//	@produce		json
//	@success		200	{json}		message
//	@failure		401	{object}	utils.ErrorResponse
//	@failure		500	{object}	utils.ErrorResponse
//	@router			/electromart/v1/customers [get]
func AllCustomers(w http.ResponseWriter, r *http.Request) {
	customersAddresses, err := dependencies.Dependencies.CustomerAddressDeps.PSQLCustomerAddress.AllCustomersAddresses(r.Context())
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	utils.JSON(w, http.StatusOK, customersAddresses)
}
