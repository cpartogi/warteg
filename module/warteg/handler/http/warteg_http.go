package http

import (
	"github.com/cpartogi/warteg/constant"
	"github.com/cpartogi/warteg/module/warteg"
	"github.com/cpartogi/warteg/pkg/utils"
	"github.com/cpartogi/warteg/schema/request"
	"github.com/labstack/echo/v4"
)

// AuthHandler  represent the httphandler for auth
type WartegHandler struct {
	wartegUsecase warteg.Usecase
}

// NewAuthHandler will initialize the contact/ resources endpoint
func NewWartegHandler(e *echo.Echo, us warteg.Usecase) {
	handler := &WartegHandler{
		wartegUsecase: us,
	}

	router := e.Group("/v1")
	router.GET("/wartegs/list", handler.WartegList)
	router.POST("/warteg", handler.WartegAdd)
	router.DELETE("/warteg/:warteg_id", handler.WartegDelete)
	router.PUT("/warteg/:warteg_id", handler.WartegUpdate)
	router.GET("/warteg/:warteg_id", handler.WartegDetail)
}

// WartegAdd godoc
// @Summary Add Warteg
// @Description Add Warteg
// @Tags Warteg
// @Accept  json
// @Produce  json
// @Param request body request.Warteg true "Request Body"
// @Success 201 {object} response.SwaggerWartegAdd
// @Failure 400 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/warteg [post]
// Wartegadd handles HTTP request for add warteg
func (h *WartegHandler) WartegAdd(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Warteg{}

	//parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, map[string]interface{}{})
	}

	//validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, map[string]interface{}{})
	}

	reg, err := h.wartegUsecase.WartegAdd(ctx, req)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.CreatedResponse(c, "Success add warteg", reg)
}

// WartegDelete godoc
// @Summary Delete Warteg
// @Description Delete Warteg
// @Tags Warteg
// @Accept  json
// @Produce  json
// @Param warteg_id path string true "Warteg_id"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/warteg/{warteg_id} [delete]
// WartegDelete handles HTTP request for delete warteg
func (h *WartegHandler) WartegDelete(c echo.Context) error {
	ctx := c.Request().Context()
	wartegId := c.Param("warteg_id")

	_, err := h.wartegUsecase.WartegDelete(ctx, wartegId)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, "Success delete warteg", map[string]interface{}{})
}

// WartegUpdate godoc
// @Summary Update Warteg
// @Description Update Warteg
// @Tags Warteg
// @Accept  json
// @Produce  json
// @Param warteg_id path string true "warteg Id"
// @Param request body request.WartegUpdate true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/warteg/{warteg_id} [put]
// WartegUpdate handles HTTP request for update warteg
func (h *WartegHandler) WartegUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	wartegId := c.Param("warteg_id")
	req := request.WartegUpdate{}

	//parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, map[string]interface{}{})
	}

	//validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, map[string]interface{}{})
	}

	_, err = h.wartegUsecase.WartegUpdate(ctx, wartegId, req)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, "Succes update warteg", map[string]interface{}{})

}

// WartegList godoc
// @Summary  Warteg list
// @Description Warteg List
// @Tags Warteg
// @Accept  json
// @Produce  json
// @Param warteg_name query string false "warteg name"
// @Success 200 {object} response.SwaggerWartegList
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/wartegs/list [get]
// WartegList handles HTTP request for warteg list
func (h *WartegHandler) WartegList(c echo.Context) error {
	ctx := c.Request().Context()
	queryValues := c.Request().URL.Query()
	warteg_name := queryValues.Get("warteg_name")

	warteg, err := h.wartegUsecase.WartegList(ctx, warteg_name)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, warteg)
}

// WartegDetail godoc
// @Summary  Warteg Detail
// @Description Warteg Detail
// @Tags Warteg
// @Accept  json
// @Produce  json
// @Param warteg_id path string true "warteg Id"
// @Success 200 {object} response.SwaggerWartegDetail
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/warteg/{warteg_id} [get]
// WartegDetail handles HTTP request for warteg detail
func (h *WartegHandler) WartegDetail(c echo.Context) error {
	ctx := c.Request().Context()
	wartegId := c.Param("warteg_id")

	wd, err := h.wartegUsecase.WartegDetail(ctx, wartegId)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, wd)
}
