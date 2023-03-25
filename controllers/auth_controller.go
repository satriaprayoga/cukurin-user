package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/cukurin-user/form"
	"github.com/satriaprayoga/cukurin-user/middlewares"
	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/pkg/logging"
	"github.com/satriaprayoga/cukurin-user/pkg/response"
	"github.com/satriaprayoga/cukurin-user/pkg/utils"
	"github.com/satriaprayoga/cukurin-user/services"
)

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(e *echo.Echo, authService services.IAuthService) {
	cont := &AuthController{
		authService: authService,
	}
	r := e.Group("/user/auth")
	r.POST("/login", cont.Login)

	l := e.Group("/user/auth/logout")
	l.Use(middlewares.JWT)
	l.POST("", cont.Logout)
}

func (c *AuthController) Logout(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		resp = response.Resp{R: e}
	)

	claims, err := form.GetClaims(e)
	if err != nil {
		return resp.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	err = c.authService.Logout(ctx, claims)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusOK, "Ok", nil)
}

func (c *AuthController) Login(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger    = logging.Logger{}
		resp      = response.Resp{R: e}
		loginForm = models.LoginForm{}
	)

	httpCode, errMsg := form.BindAndValid(e, &loginForm)
	logger.Info(utils.Stringify(loginForm))
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", errMsg), nil)
	}

	out, err := c.authService.Login(ctx, &loginForm)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusOK, "Ok", out)
}
