package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/cukurin-user/form"
	"github.com/satriaprayoga/cukurin-user/middlewares"
	"github.com/satriaprayoga/cukurin-user/pkg/logging"
	"github.com/satriaprayoga/cukurin-user/pkg/response"
	"github.com/satriaprayoga/cukurin-user/services"
)

type KUserController struct {
	kUserService services.IKUserService
}

func NewKUserController(e *echo.Echo, k services.IKUserService) {
	controller := &KUserController{
		kUserService: k,
	}
	r := e.Group("/user/user")
	r.Use(middlewares.JWT)
	r.GET("/account", controller.GetAccount)
}

func (k *KUserController) GetAccount(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		resp   = response.Resp{R: e}
	)
	claims, err := form.GetClaims(e)
	logger.Info("%v", claims)
	if err != nil {
		return resp.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	data, err := k.kUserService.GetDataBy(ctx, claims, claims.UserID)
	if err != nil {
		return resp.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusOK, "Ok", data)

}
