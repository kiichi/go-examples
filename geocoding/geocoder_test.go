package geocoding

import (
	"testing"
)

// read Schools, do geo coding, save it
func TestGeocode(t *testing.T) {
	schools := ReadSchools()
	schools_geocoded := Geocode(schools)
	WriteSchools(schools_geocoded)
}
