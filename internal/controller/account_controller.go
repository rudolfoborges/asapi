package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/rudolfoborges/asapi/internal/service"
)

type CreateAccountRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountController struct {
	accountService *service.AccountService
}

func NewAccountController(accountService *service.AccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

func (c *AccountController) CreateAccount(ctx echo.Context) error {
	var request CreateAccountRequest

	err := ctx.Bind(&request)
	if err != nil {
		return ctx.JSON(400, map[string]string{"error": "Invalid request"})
	}

	err = c.accountService.Create(ctx.Request().Context(), service.CreateAccountInput{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return ctx.JSON(500, map[string]string{"error": "Failed to create account"})
	}

	return ctx.JSON(201, map[string]string{"message": "Account created successfully"})
}
