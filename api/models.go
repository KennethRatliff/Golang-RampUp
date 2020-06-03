package api

type (
	// Location is the latitude and longitude of a driver
	Location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	}

	// Payload contains the needed driver data
	Payload struct {
		Timestamp int64    `json:"timestamp"`
		DriverID  int      `json:"driver_id"`
		Location  Location `json:"location"`
	}

	// DefaultResponse is used for responding to non driver and non nearest driver requests
	DefaultResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	// DriverResponse is the response format for responding to a driver request
	DriverResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Driver  int    `json:"driver"`
	}

	// NearestDriverResponse is the response format for a nearest driver request
	NearestDriverResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Drivers []int  `json:"drivers"`
	}
)
