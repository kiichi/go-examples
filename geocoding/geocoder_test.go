package geocoding

import (
	"testing"
)

func TestGeocode(t *testing.T) {
	schools := ReadSchools()
	Geocode(schools)
}
