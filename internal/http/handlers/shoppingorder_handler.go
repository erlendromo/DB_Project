package handlers

import (
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/http/middlewares"
	"DB_Project/internal/utils"
	"errors"
	"net/http"
	"strconv"
)

// CreateOrder handler
//
//	@title			CreateOrder
//	@Summary		Create an order
//	@description	Creates an order for a logged in customer
//	@tags			shoppingorder
//	@produce		json
//	@success		201	{object}	shoppingorderdomain.ShoppingOrderResponse
//	@failure		400	{object}	utils.ErrorResponse
//	@failure		401	{object}	utils.ErrorResponse
//	@router			/electromart/v1/checkout [post]
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if len(Cart) == 0 {
		utils.ERROR(w, http.StatusBadRequest, errors.New("cart is empty"))
		return
	}

	sessionData, statuscode, err := middlewares.GetUserFromSession(r)
	if err != nil {
		utils.ERROR(w, statuscode, err)
		return
	}

	orderResponse, err := dependencies.Dependencies.ShoppingOrderDeps.PSQLShoppingOrder.CreateOrder(r.Context(), sessionData.ID, Cart)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	ResetCart()

	utils.JSON(w, http.StatusCreated, orderResponse)
}

// GetShoppingOrderByID handler
//
//	@title			GetShoppingOrderByID
//	@Summary		Get an order by ID
//	@description	Gets an order by ID of a logged in customer
//	@tags			shoppingorder
//	@produce		json
//	@Param			orderID	path		int	true	"Order ID"
//	@success		200		{object}	shoppingorderdomain.ShoppingOrderResponse
//	@failure		400		{object}	utils.ErrorResponse
//	@failure		401		{object}	utils.ErrorResponse
//	@router			/electromart/v1/orders/{orderID} [get]
func GetShoppingOrderByID(w http.ResponseWriter, r *http.Request) {
	sessionData, statuscode, err := middlewares.GetUserFromSession(r)
	if err != nil {
		utils.ERROR(w, statuscode, err)
		return
	}

	shoppingOrderID := r.PathValue("orderID")

	shoppingOrderIDInt, err := strconv.Atoi(shoppingOrderID)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, errors.New("invalid orderID"))
		return
	}

	orderResponse, err := dependencies.Dependencies.ShoppingOrderDeps.PSQLShoppingOrder.GetOrderByID(r.Context(), sessionData.ID, shoppingOrderIDInt)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	utils.JSON(w, http.StatusOK, orderResponse)
}

// GetShoppingOrders handler
//
//	@title			GetShoppingOrders
//	@Summary		Get all orders
//	@description	Gets all orders of a logged in customer
//	@tags			shoppingorder
//	@produce		json
//	@success		200	{object}	[]shoppingorderdomain.ShoppingOrderResponse
//	@failure		400	{object}	utils.ErrorResponse
//	@failure		401	{object}	utils.ErrorResponse
//	@router			/electromart/v1/orders [get]
func GetShoppingOrders(w http.ResponseWriter, r *http.Request) {
	sessionData, statuscode, err := middlewares.GetUserFromSession(r)
	if err != nil {
		utils.ERROR(w, statuscode, err)
		return
	}

	orders, err := dependencies.Dependencies.ShoppingOrderDeps.PSQLShoppingOrder.GetOrders(r.Context(), sessionData.ID)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	utils.JSON(w, http.StatusOK, orders)
}
