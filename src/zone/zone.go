package zone

import (
	"encoding/xml"
	"strconv"
)

// Territories contains Territory's within defining spawns
type Territories struct {
	XMLName     xml.Name    `xml:"territory-type"`
	Territories []Territory `xml:"territory"`
}

// Territory contains a list of Zones
type Territory struct {
	XMLName xml.Name `xml:"territory"`
	Zones   []Zone   `xml:"zone"`
	Color   string   `xml:"color,attr"`
}

// Zone represents a spawnpoint and rates of an entity 
type Zone struct {
	XMLName xml.Name `xml:"zone"`
	Name    string   `xml:"name,attr"`
	Smin    string   `xml:"smin,attr"`
	Smax    string   `xml:"smax,attr"`
	Dmin    string   `xml:"dmin,attr"`
	Dmax    string   `xml:"dmax,attr"`
	X       string   `xml:"x,attr"`
	Z       string   `xml:"z,attr"`
	R       string   `xml:"r,attr"`
}

// MultiplyZone values by provided amounts
func (zone *Zone) MultiplyZone(multiplier float64, radiusMultiplier float64, affectMin bool) {
	if affectMin {
		// smin
		smin, _ := strconv.Atoi(zone.Smin)
		zone.Smin = strconv.Itoa(int(float64(smin) * multiplier))

		// dmin
		dmin, _ := strconv.Atoi(zone.Dmin)
		zone.Dmin = strconv.Itoa(int(float64(dmin) * multiplier))
	}
	// smax
	smax, _ := strconv.Atoi(zone.Smax)
	zone.Smax = strconv.Itoa(int(float64(smax) * multiplier))
	// dmax
	dmax, _ := strconv.Atoi(zone.Dmax)
	zone.Dmax = strconv.Itoa(int(float64(dmax) * multiplier))
	// r
	r, _ := strconv.Atoi(zone.R)
	zone.R = strconv.Itoa(int(float64(r) * radiusMultiplier))
}
