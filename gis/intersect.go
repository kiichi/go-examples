package gis

import (
	"fmt"
	"github.com/paulmach/go.geo"
)

func IntersectExample() {
	// lng,lat
	line1 := geo.NewLine(geo.NewPoint(-73.942799, 40.7459951), geo.NewPoint(-73.937349, 40.748783))
	line2 := geo.NewLine(geo.NewPoint(-73.944494, 40.751250), geo.NewPoint(-73.937010, 40.743877))
	// a line far away from those above lines
	line3 := geo.NewLine(geo.NewPoint(-73.949167, 40.739304), geo.NewPoint(-73.949167, 40.739304))

	if line1.Intersects(line2) {
		fmt.Println("ok. l1 and l2 should be intersecting")
	} else {
		fmt.Println("hey... this is wrong. somehow l1 and l2 are NOT intersecting")

	}

	if line1.Intersects(line3) {
		fmt.Println("wait ... somehow l1 and l2 are intersecting")
	} else {
		fmt.Println("this is correct. l1 and l2 should not be intersecting")

	}

}
