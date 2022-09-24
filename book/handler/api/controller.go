package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/Jiran03/agmc/task/day5/book/domain"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewBookHandler(service domain.Service) BookHandler {
	return BookHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (bh BookHandler) Create(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := bh.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := bh.service.InsertData(toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
	})
}

func (bh BookHandler) GetAll(ctx echo.Context) error {
	bookRes, err := bh.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	bookObj := []ResponseJSON{}

	for _, value := range bookRes {
		bookObj = append(bookObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    bookObj,
	})
}

func (bh BookHandler) GetByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	bookRes, err := bh.service.GetByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	bookObj := fromDomain(bookRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    bookObj,
	})
}

func (bh BookHandler) Update(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	bookRes, err := bh.service.UpdateData(id, toDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	bookObj := fromDomain(bookRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    bookObj,
	})
}

func (bh BookHandler) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := bh.service.DeleteData(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
	})
}
