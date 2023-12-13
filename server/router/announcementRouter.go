package router

import (
	"server/internal/controller"
	"server/internal/repository"
	"server/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AnnouncementRouter(e *echo.Echo, mainUrl string, db *gorm.DB, validate *validator.Validate) {
	announcementRepository := repository.NewAnnouncementRepository()
	announcementService := service.NewAnnouncementService(db, validate, announcementRepository)
	announcementController := controller.NewAnnouncementController(announcementService)

	g := e.Group(mainUrl + "/announcements")
	g.GET("", announcementController.GetAll)
	g.GET("/:id", announcementController.GetById)
	g.POST("", announcementController.Save)
	g.PUT("/:id", announcementController.Update)
	g.DELETE("/:id", announcementController.Delete)
}
