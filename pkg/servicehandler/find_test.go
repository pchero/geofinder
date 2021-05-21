package servicehandler

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestGeoFind(t *testing.T) {

	// create mock
	mc := gomock.NewController(t)
	defer mc.Finish()

	// mockSvc := servicehandler.NewMockServiceHandler(mc)
	h := NewServiceHandler("../../GeoLite2-City.mmdb")

	type test struct {
		name      string
		address   string
		countries []string
		epxectRes bool
	}

	tests := []test{
		{
			"normal 1 country",
			"81.2.69.142",
			[]string{"GB"},
			true,
		},
		{
			"normal 2 countires",
			"211.200.20.28",
			[]string{"KR", "GB"},
			true,
		},
		{
			"empty countires",
			"211.200.20.28",
			[]string{},
			false,
		},
		{
			"not in the list",
			"211.200.20.28",
			[]string{"GB", "US"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, err := h.GeoFind(tt.address, tt.countries)
			if err != nil {
				t.Errorf("Wrong match. expect: ok, gor: %v", err)
			}

			if res != tt.epxectRes {
				t.Errorf("Wrong match. expect: %v, got: %v", tt.epxectRes, res)
			}
		})
	}
}
