package exception

import (
	"errors"
	"net/http"
	"server/model/web"

	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, c echo.Context) {
	var res web.WebErrorResponse

	if castedErr, ok := err.(*BadRequestError); ok {
		res.Code = http.StatusBadRequest
		res.Status = "BAD REQUEST"
		res.Message = castedErr.Error()
	} else if castedErr, ok := err.(*NotFoundError); ok {
		res.Code = http.StatusNotFound
		res.Status = "NOT FOUND"
		res.Message = castedErr.Error()
	} else if ok := errors.Is(err, echo.ErrNotFound); ok {
		res.Code = http.StatusNotFound
		res.Status = "NOT FOUND"
		res.Message = err.Error()
	} else {
		res.Code = http.StatusInternalServerError
		res.Status = "INTERNAL SERVER ERROR"
		res.Message = err.Error()
	}

	c.JSON(res.Code, res)
}
