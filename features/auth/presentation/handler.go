package presentation

import (
	"net/http"
	"project/e-comerce/features/auth"
	"project/e-comerce/features/auth/presentation/request"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{
	userBusiness auth.Business
}

func NewUserHandler(business auth.Business) *UserHandler{
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler)Login(c echo.Context) error{
	reqBody := request.User{}
	errBind := c.Bind(&reqBody)
	if errBind != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get bind data",
		})
	}

	authCore := request.ToCore(reqBody)
	result, name, err:= h.userBusiness.Login(authCore)
	
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get token data"+err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
		"name" : name,
		"token" : result,
	})
}

