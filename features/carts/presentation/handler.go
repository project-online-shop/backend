package presentation

import (
	"errors"
	"net/http"
	"project/e-comerce/features/carts"
	_requestCart "project/e-comerce/features/carts/presentation/request"
	_responseCart "project/e-comerce/features/carts/presentation/response"
	"project/e-comerce/helper"
	"project/e-comerce/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartBusiness carts.Business
}

func NewCartHandler(business carts.Business) *CartHandler {
	return &CartHandler{
		cartBusiness: business,
	}
}

func (h *CartHandler) GetAll(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to get user id"))
	}
	UserId := userID_token
	result, err := h.cartBusiness.GetAllCart(UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(" failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success get data", _responseCart.FromCoreList(result)))
}

func (h *CartHandler) AddCart(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to get user id"))
	}

	cartData := _requestCart.Cart{}
	errBind := c.Bind(&cartData)

	if errBind != nil {
		return errors.New("failed to bind data")
	}

	cartCore := _requestCart.ToCore(cartData)
	cartCore.UserID = userID_token
	cartCore.Product.ID = cartData.ProductID

	result, err := h.cartBusiness.AddCart(cartCore)

	if err != nil {
		return errors.New("failed to insert data")
	}

	if result == 0 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to add data"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
}

func (h *CartHandler) Update(c echo.Context) error {
	idCart, _ := strconv.Atoi(c.Param("id"))
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to get user id"))
	}

	cartData := _requestCart.Cart{}
	errBind := c.Bind(&cartData)

	if errBind != nil {
		return errors.New("failed to bind data")
	}

	cartCore := _requestCart.ToUpdateCore(cartData)
	cartCore.UserID = userID_token

	result, err := h.cartBusiness.UpdateCart(cartCore, idCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to update cart"))
	}
	if result == 0 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to update cart"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to update data"))
}

func (h *CartHandler) Destroy(c echo.Context) error {
	idCart, _ := strconv.Atoi(c.Param("id"))
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to get user id"))
	}
	UserId := userID_token

	result, err := h.cartBusiness.DestroyCart(UserId, idCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error response to delete cart"))
	}
	if result == 0 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to delete cart"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to delete data"))
}
