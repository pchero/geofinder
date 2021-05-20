package geohandler

// ServiceHandler is interface for service handle
type GeoHandler interface {
	FindCountry(address string) (string, error)
}

type geoHandler struct {
}

// NewGeoHandler return ServiceHandler interface
func NewGeoHandler() GeoHandler {
	return &geoHandler{}
}
