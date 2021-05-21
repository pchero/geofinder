package geohandler

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFindCountry(t *testing.T) {

	// create mock
	mc := gomock.NewController(t)
	defer mc.Finish()

	// mockSvc := servicehandler.NewMockServiceHandler(mc)
	h := NewGeoHandler("../../GeoLite2-City.mmdb")

	type test struct {
		name      string
		address   string
		expectRes string
	}

	tests := []test{
		{
			"normal GB",
			"81.2.69.142",
			"GB",
		},
		{
			"normal KR",
			"211.200.20.28",
			"KR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, err := h.FindCountry(tt.address)
			if err != nil {
				t.Errorf("Wrong match. expect: ok, gor: %v", err)
			}

			if res != tt.expectRes {
				t.Errorf("Wrong match. expect: %s, got: %s", tt.expectRes, res)
			}
		})
	}
}
