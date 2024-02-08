package handler

import (
	"net/http"
	customers "store/features/customers/entity"

	"github.com/labstack/echo/v4"
)

type CustomerController struct {
	customerUseCase customers.CustomerUseCaseInterface
}

func New(customerUC customers.CustomerUseCaseInterface) *CustomerController {
	return &CustomerController{
		customerUseCase: customerUC,
	}
}

func (handler *CustomerController) Login(c echo.Context) error {
	input := new(CustomerRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := customers.CustomerCore{
		Email:    input.Email,
		Password: input.Password,
	}

	loginData, token, err := handler.customerUseCase.Login(data.Email, data.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Login failed. Email or password incorrect.",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "login success",
		"email":   loginData.Email,
		"token":   token,
	})
}

func (handler *CustomerController) Register(c echo.Context) error {
	input := new(CustomerRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := customers.CustomerCore{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Address:  input.Address,
	}

	_, errusers := handler.customerUseCase.Register(data)
	if errusers != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create user",
			"error":   errusers.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create user",
	})
}
