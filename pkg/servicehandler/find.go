package servicehandler

import (
	"github.com/sirupsen/logrus"
)

// Find is a service handler for find geo.
func (h *serviceHandler) GeoFind(address string, countries []string) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"address":   address,
		"countries": countries,
	})

	country, err := h.geoHandler.FindCountry(address)
	if err != nil {
		log.Errorf("Could not get geo info. err: %v", err)
		return false, err
	}

	// check the country is in the list
	for _, c := range countries {
		if c == country {
			return true, nil
		}
	}

	// not listed
	return false, nil
}
