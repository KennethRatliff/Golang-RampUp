package storage

type (
	// Location used for storing driver's location
	Location struct {
		Lat float64
		Lon float64
	}
	// Driver model to store driver data
	Driver struct {
		ID           int
		LastLocation Location
	}
)

// DriverStorage is main storage for our project
type DriverStorage struct {
	drivers map[int]*Driver
}

// New creates new instance of DriverStorage
func New() *DriverStorage {
	d := &DriverStorage{}
	d.drivers = make(map[int]*Driver)
	return d
}

// Set sets driver to storage by key
func (d *DriverStorage) Set(key int, driver *Driver) {
	d.drivers[key] = driver
}

// Delete removes driver from storage by key
func (d *DriverStorage) Delete(key int) error {
	return nil
}

// Get gets driver from storage and an error if nothing found
func (d *DriverStorage) Get(key int) (*Driver, error) {
	return nil, nil
}

// Nearest returns nearest drivers by locaion
func (d *DriverStorage) Nearest(lat, lon float64) ([]*Driver, error) {
	return nil, nil
}
