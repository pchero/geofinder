package geohandler

//go:generate go run -mod=mod github.com/golang/mock/mockgen -package geohandler -destination ./mock_geohandler_geohandler.go -source main.go -build_flags=-mod=mod

import (
	"log"

	"github.com/oschwald/geoip2-golang"
)

// ServiceHandler is interface for service handle
type GeoHandler interface {
	FindCountry(address string) (string, error)
}

type geoHandler struct {
	db *geoip2.Reader
}

// NewGeoHandler return ServiceHandler interface
func NewGeoHandler(dbfile string) GeoHandler {

	db, err := geoip2.Open(dbfile)
	if err != nil {
		log.Fatal(err)
	}

	return &geoHandler{
		db: db,
	}
}
