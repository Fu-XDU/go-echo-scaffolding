package controller

import (
	"server/model/base"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Info(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, base.NewResponse(base.SUCCESS))
}
