package handlers

import (
	"DB_Project/internal/utils"
	"errors"
	"net/http"
	"strconv"
)

var Cart map[int]int

// AddToCart handler
//
//	@title			AddToCart
//	@Summary		Add a product to cart
//	@description	Adds a product to cart
//	@tags			cart
//	@param			productID	path	int	true	"Product ID"
//	@param			quantity	query	int	false	"Quantity"
//	@produce		json
//	@success		200	{object}	CartResponse
//	@failure		400	{object}	utils.ErrorResponse
//	@router			/electromart/v1/cart/{productID} [post]
func AddToCart(w http.ResponseWriter, r *http.Request) {
	if Cart == nil {
		ResetCart()
	}

	productID := r.PathValue("productID")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	quantity := r.URL.Query().Get("quantity")
	if len(quantity) == 0 {
		Cart[productIDInt] = 1
	}

	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if quantityInt < 1 {
		utils.ERROR(w, http.StatusBadRequest, errors.New("invalid quantity"))
		return
	}

	Cart[productIDInt] = quantityInt

	utils.JSON(w, http.StatusOK, cartToJSON())
}

// GetCart handler
//
//	@title			GetCart
//	@Summary		Get the cart
//	@description	Returns the cart
//	@tags			cart
//	@produce		json
//	@success		200	{object}	CartResponse
//	@router			/electromart/v1/cart [get]
func GetCart(w http.ResponseWriter, r *http.Request) {
	if Cart == nil {
		ResetCart()
	}
	utils.JSON(w, http.StatusOK, cartToJSON())
}

func ResetCart() {
	Cart = make(map[int]int, 0)
}

type CartResponse struct {
	CartItems []CartItem `json:"cart_items"`
}

type CartItem struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

func cartToJSON() CartResponse {
	cartResponse := CartResponse{}
	for productID, quantity := range Cart {
		cartItem := CartItem{
			ProductID: productID,
			Quantity:  quantity,
		}
		cartResponse.CartItems = append(cartResponse.CartItems, cartItem)
	}

	return cartResponse
}
