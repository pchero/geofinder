package servicehandler

//go:generate go run -mod=mod github.com/golang/mock/mockgen -package servicehandler -destination ./mock_servicehandler_servicehandler.go -source main.go -build_flags=-mod=mod

import (
	"strings"
	"time"

	"github.com/pchero/geofinder/pkg/geohandler"
)

// ServiceHandler is interface for service handle
type ServiceHandler interface {
	GeoFind(address string, countries []string) (bool, error)
}

type serviceHandler struct {
	geoHandler geohandler.GeoHandler
}

// NewServiceHandler return ServiceHandler interface
func NewServiceHandler(geodb string) ServiceHandler {
	geoHandler := geohandler.NewGeoHandler(geodb)

	return &serviceHandler{
		geoHandler: geoHandler,
	}
}

// getCurTime return current utc time string
func getCurTime() string {
	now := time.Now().UTC().String()
	res := strings.TrimSuffix(now, " +0000 UTC")

	return res
}
