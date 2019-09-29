package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/fatih/color"
)

// Territories in our file
type Territories struct {
	XMLName     xml.Name    `xml:"territory-type"`
	Territories []Territory `xml:"territory"`
}

// Territory contain
type Territory struct {
	XMLName xml.Name `xml:"territory"`
	Zones   []Zone   `xml:"zone"`
	Color   string   `xml:"color,attr"`
}

// Zone event
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

// Flags available for change
type Flags struct {
	AffectMin           *bool
	Radius              *int
	InfectedGlobal      *int
	InfectedArmy        *int
	InfectedVillage     *int
	InfectedMedic       *int
	InfectedPolice      *int
	InfectedReligious   *int
	InfectedIndustrial  *int
	InfectedFirefighter *int
	InfectedCity        *int
	InfectedSolitude    *int
}

func buildMultipliedZone(multiplier int, radiusMultiplier int, affectMin bool, zone Zone) Zone {
	newZone := zone
	if affectMin {
		// smin
		smin, _ := strconv.Atoi(zone.Smin)
		newZone.Smin = strconv.Itoa(smin * multiplier)

		// dmin
		dmin, _ := strconv.Atoi(zone.Dmin)
		newZone.Dmin = strconv.Itoa(dmin * multiplier)
	}
	// smax
	smax, _ := strconv.Atoi(zone.Smax)
	newZone.Smax = strconv.Itoa(smax * multiplier)
	// dmax
	dmax, _ := strconv.Atoi(zone.Dmax)
	newZone.Dmax = strconv.Itoa(dmax * multiplier)
	// r
	r, _ := strconv.Atoi(zone.R)
	newZone.R = strconv.Itoa(r * radiusMultiplier)
	return newZone
}

func main() {
	// Flags
	flags := Flags{
		AffectMin:           flag.Bool("AffectMin", false, "Also multiplies minimum zombie spawn rate if true"),
		Radius:              flag.Int("Radius", 1, "Infected radius spawn multiplier."),
		InfectedGlobal:      flag.Int("InfectedGlobal", 1, "InfectedGlobal multiplier amount (Real Number)."),
		InfectedArmy:        flag.Int("InfectedArmy", 1, "InfectedArmy multiplier amount (Real Number)."),
		InfectedVillage:     flag.Int("InfectedVillage", 1, "InfectedVillage multiplier amount (Real Number)."),
		InfectedMedic:       flag.Int("InfectedMedic", 1, "InfectedMedic multiplier amount (Real Number)."),
		InfectedPolice:      flag.Int("InfectedPolice", 1, "InfectedPolice multiplier amount (Real Number)."),
		InfectedReligious:   flag.Int("InfectedReligious", 1, "InfectedReligious multiplier amount (Real Number)."),
		InfectedIndustrial:  flag.Int("InfectedIndustrial", 1, "InfectedIndustrial multiplier amount (Real Number)."),
		InfectedFirefighter: flag.Int("InfectedFirefighter", 1, "InfectedFirefighter multiplier amount (Real Number)."),
		InfectedCity:        flag.Int("InfectedCity", 1, "InfectedCity multiplier amount (Real Number)."),
		InfectedSolitude:    flag.Int("InfectedSolitude", 1, "InfectedSolitude multiplier amount (Real Number)."),
	}

	// Parse flags
	flag.Parse()

	color.Blue("Sepcial Flags:")
	fmt.Printf(`AffectMin		- %v`+"\n\n", *flags.AffectMin)
	color.Blue("Multipliers:")
	fmt.Printf(`Radius			- %v
Global			- %v
InfectedArmy 		- %v
InfectedVillage 	- %v
InfectedMedic 		- %v
InfectedPolice 		- %v
InfectedReligious 	- %v
InfectedIndustrial 	- %v
InfectedFirefighter	- %v
InfectedCity		- %v
InfectedSolitude	- %v`+"\n\n",
		*flags.Radius,
		*flags.InfectedGlobal,
		*flags.InfectedArmy,
		*flags.InfectedVillage,
		*flags.InfectedMedic,
		*flags.InfectedPolice,
		*flags.InfectedReligious,
		*flags.InfectedIndustrial,
		*flags.InfectedFirefighter,
		*flags.InfectedCity,
		*flags.InfectedSolitude,
	)

	// Read XML
	zombieTerritoriesXML, err := os.Open("xml/zombie_territories.base.xml")
	// Error reading or finding XML file
	if err != nil {
		fmt.Println(err)
	}

	color.Blue("Opened base Zombie Territories XML...\n")
	// Keep file open
	defer zombieTerritoriesXML.Close()

	byteValue, _ := ioutil.ReadAll(zombieTerritoriesXML)

	var territories Territories
	xml.Unmarshal(byteValue, &territories)

	// Scan territories
	for i := 0; i < len(territories.Territories); i++ {
		fmt.Printf("Territoriy %v found (%v zones)...\n",
			territories.Territories[i].Color,
			len(territories.Territories[i].Zones),
		)

		// Scan zones
		selectedZones := territories.Territories[i].Zones
		for j := 0; j < len(selectedZones); j++ {
			// Activated multipliers
			territories.Territories[i].Zones[j] = buildMultipliedZone(
				*flags.InfectedArmy+*flags.InfectedGlobal-1,
				*flags.Radius,
				*flags.AffectMin,
				selectedZones[j],
			)
		}
		fmt.Printf("Multipliers applied to %v zones. \n", len(selectedZones))
	}

	file, _ := xml.MarshalIndent(territories, "", "	")
	_ = ioutil.WriteFile("xml/zombie_territories.xml", file, 0644)
	color.Green("\n\nWrote `zombie_territories.xml` to the `xml/` directory. Upload to your server to modify spawns")
}
