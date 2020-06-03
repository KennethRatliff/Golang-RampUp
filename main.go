package main

import (
	"log"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	g := e.Group("/api")
	g.POST("/driver/", addDriver)
	g.GET("/driver/:id", getDriver)
	g.DELETE("/driver/:id", deleteDriver)
	g.GET("/driver/:lat/:lon/nearest", nearestDrivers)
	log.Fatal(e.Start(":9111"))
}

type (
	// Location format
	Location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	}
	// Payload format
	Payload struct {
		Timestamp int64    `json:"timestamp"`
		DriverID  int      `json:"driver_id"`
		Location  Location `json:"location"`
	}
)

type (
	// DefaultResponse used for returning default response
	DefaultResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	// DriverResponse returns driver
	DriverResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Driver  int    `json:"driver"`
	}
	// NearestDriverResponse returns nearest drivers
	NearestDriverResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Drivers []int  `json:"drivers"`
	}
)

func addDriver(c echo.Context) error {
	return nil
}
func getDriver(c echo.Context) error {
	return nil
}
func deleteDriver(c echo.Context) error {
	return nil
}
func nearestDrivers(c echo.Context) error {
	return nil
}
