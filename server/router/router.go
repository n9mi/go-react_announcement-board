package router

import (
	"server/exception"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitializeEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.HTTPErrorHandler = exception.CustomErrorHandler

	return e
}

func AssignRouter(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	mainUrl := "/api"

	AnnouncementRouter(e, mainUrl, db, validate)
}
