package handlers

import (
	"CalculatorAppBackend/internal/calculationService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalculationHandler struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(s calculationService.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: s}
}

func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
	}
	return c.JSON(http.StatusOK, calculations)
}

func (h *CalculationHandler) PostCalculations(c echo.Context) error {
	var req calculationService.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.CreateCalculation(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create calculation"})
	}

	return c.JSON(http.StatusCreated, calc)
}

func (h *CalculationHandler) PatchCalculations(c echo.Context) error {
	id := c.Param("id")

	var req calculationService.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	updatedCalc, err := h.service.UpdateCalculation(id, req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update calculation"})
	}
	return c.JSON(http.StatusOK, updatedCalc)
}

func (h *CalculationHandler) DeleteCalculations(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete calculation"})
	}
	return c.NoContent(http.StatusNoContent)
}
