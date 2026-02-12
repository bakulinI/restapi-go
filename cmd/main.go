package main

import (
	"CalculatorAppBackend/internal/calculationService"
	"CalculatorAppBackend/internal/db"
	"CalculatorAppBackend/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)

	e.Start("localhost:8080")
}
