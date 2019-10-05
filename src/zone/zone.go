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

func convertMultiply(a string, multiplier float64) string {
	b, _ := strconv.Atoi(a)
	return strconv.Itoa(int(float64(b) * multiplier))
}

// MultiplyZone values by provided amounts
func (zone *Zone) MultiplyZone(multiplier float64, radiusMultiplier float64, affectMin bool) {
	if affectMin {
		// smin
		zone.Smin = convertMultiply(zone.Smin, multiplier)
		// dmin
		zone.Dmin = convertMultiply(zone.Dmin, multiplier)
	}
	// smax
	zone.Smax = convertMultiply(zone.Smax, multiplier)
	// dmax
	zone.Dmax = convertMultiply(zone.Dmax, multiplier)
	// r
	zone.R = convertMultiply(zone.R, multiplier)
}
