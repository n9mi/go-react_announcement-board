package controller

import (
	"net/http"
	"server/exception"
	"server/internal/service"
	"server/model/web"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AnnouncementController interface {
	GetAll(c echo.Context) error
	Save(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type announcementControllerImpl struct {
	AnnouncementService service.AnnouncementService
}

func NewAnnouncementController(announcementService service.AnnouncementService) *announcementControllerImpl {
	return &announcementControllerImpl{
		AnnouncementService: announcementService,
	}
}

func (ct *announcementControllerImpl) GetAll(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	anns, err := ct.AnnouncementService.FindAll(c.Request().Context(), page, pageSize)
	if err != nil {
		return err
	}

	res := web.WebSuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   anns,
	}
	return c.JSON(res.Code, res)
}

func (ct *announcementControllerImpl) Save(c echo.Context) error {
	req := new(web.Announcement)

	if err := c.Bind(req); err != nil {
		return &exception.BadRequestError{Message: err.Error()}
	}

	ann, err := ct.AnnouncementService.Create(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	res := web.WebSuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ann,
	}
	return c.JSON(res.Code, res)
}

func (ct *announcementControllerImpl) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return &exception.NotFoundError{Entity: "announcement"}
	}

	ann, err := ct.AnnouncementService.FindById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	res := web.WebSuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ann,
	}
	return c.JSON(res.Code, res)
}

func (ct *announcementControllerImpl) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return &exception.NotFoundError{Entity: "announcement"}
	}

	req := new(web.Announcement)
	if err := c.Bind(req); err != nil {
		return &exception.BadRequestError{Message: err.Error()}
	}

	req.ID = uint64(id)
	ann, err := ct.AnnouncementService.Update(c.Request().Context(), *req)
	if err != nil {
		return err
	}

	res := web.WebSuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ann,
	}
	return c.JSON(res.Code, res)
}

func (ct *announcementControllerImpl) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return &exception.NotFoundError{Entity: "announcement"}
	}

	if err := ct.AnnouncementService.Delete(c.Request().Context(), id); err != nil {
		return err
	}

	res := web.WebSuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}
	return c.JSON(res.Code, res)
}
