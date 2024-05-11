package handlers

import (
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/http/middlewares"
	"DB_Project/internal/utils"
	"errors"
	"net/http"
)

// CreateOrder handler
//
//	@title			CreateOrder
//	@Summary		Create an order
//	@description	Creates an order
//	@tags			shoppingorder
//	@produce		json
//	@success		201	{object}	shoppingorderdomain.ShoppingOrderResponse
//	@failure		401	{object}	utils.ErrorResponse
//	@failure		500	{object}	utils.ErrorResponse
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
		utils.ERROR(w, http.StatusInternalServerError, utils.NewInternalServerError(err))
		return
	}

	ResetCart()

	utils.JSON(w, http.StatusCreated, orderResponse)
}
