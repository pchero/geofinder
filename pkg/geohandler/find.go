package geohandler

import (
	"net"

	"github.com/sirupsen/logrus"
)

func (h *geoHandler) FindCountry(address string) (string, error) {

	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(address)
	record, err := h.db.City(ip)
	if err != nil {
		logrus.Errorf("Could not parse the address correctly. err: %v", err)
		return "", err
	}

	return record.Country.IsoCode, nil
}
