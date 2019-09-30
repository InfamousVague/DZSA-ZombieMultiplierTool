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
	AdditionalSpawns    *bool
	AffectMin           *bool
	Radius              *float64
	InfectedGlobal      *float64
	InfectedArmy        *float64
	InfectedVillage     *float64
	InfectedMedic       *float64
	InfectedPolice      *float64
	InfectedReligious   *float64
	InfectedIndustrial  *float64
	InfectedFirefighter *float64
	InfectedCity        *float64
	InfectedSolitude    *float64
}

func buildMultipliedZone(multiplier float64, radiusMultiplier float64, affectMin bool, zone Zone) Zone {
	newZone := zone
	if affectMin {
		// smin
		smin, _ := strconv.Atoi(zone.Smin)
		newZone.Smin = strconv.Itoa(int(float64(smin) * multiplier))

		// dmin
		dmin, _ := strconv.Atoi(zone.Dmin)
		newZone.Dmin = strconv.Itoa(int(float64(dmin) * multiplier))
	}
	// smax
	smax, _ := strconv.Atoi(zone.Smax)
	newZone.Smax = strconv.Itoa(int(float64(smax) * multiplier))
	// dmax
	dmax, _ := strconv.Atoi(zone.Dmax)
	newZone.Dmax = strconv.Itoa(int(float64(dmax) * multiplier))
	// r
	r, _ := strconv.Atoi(zone.R)
	newZone.R = strconv.Itoa(int(float64(r) * radiusMultiplier))
	return newZone
}

func main() {
	// Flags
	flags := Flags{
		AdditionalSpawns:    flag.Bool("AdditionalSpawns", false, "Load in additional spawns xml if true"),
		AffectMin:           flag.Bool("AffectMin", false, "Also multiplies minimum zombie spawn rate if true"),
		Radius:              flag.Float64("Radius", 1, "Infected radius spawn multiplier."),
		InfectedGlobal:      flag.Float64("InfectedGlobal", 1, "InfectedGlobal multiplier amount (Real Number)."),
		InfectedArmy:        flag.Float64("InfectedArmy", 1, "InfectedArmy multiplier amount (Real Number)."),
		InfectedVillage:     flag.Float64("InfectedVillage", 1, "InfectedVillage multiplier amount (Real Number)."),
		InfectedMedic:       flag.Float64("InfectedMedic", 1, "InfectedMedic multiplier amount (Real Number)."),
		InfectedPolice:      flag.Float64("InfectedPolice", 1, "InfectedPolice multiplier amount (Real Number)."),
		InfectedReligious:   flag.Float64("InfectedReligious", 1, "InfectedReligious multiplier amount (Real Number)."),
		InfectedIndustrial:  flag.Float64("InfectedIndustrial", 1, "InfectedIndustrial multiplier amount (Real Number)."),
		InfectedFirefighter: flag.Float64("InfectedFirefighter", 1, "InfectedFirefighter multiplier amount (Real Number)."),
		InfectedCity:        flag.Float64("InfectedCity", 1, "InfectedCity multiplier amount (Real Number)."),
		InfectedSolitude:    flag.Float64("InfectedSolitude", 1, "InfectedSolitude multiplier amount (Real Number)."),
	}

	// Parse flags
	flag.Parse()

	color.Blue("Sepcial Flags:")
	fmt.Printf(`AffectMin		- %v
AdditionalSpawns	- %v`+"\n\n",
		*flags.AffectMin,
		*flags.AdditionalSpawns,
	)
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
		color.White("Territoriy %v found (%v zones)...\n",
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
		color.White("Multipliers applied to %v zones. \n", len(selectedZones))
	}

	// Load additional zombie spawns
	if *flags.AdditionalSpawns {
		color.Blue("Applying additional spawnpoints...")
		additionalTerritoriesXML, err := os.Open("xml/additional_spawns.xml")
		if err != nil {
			fmt.Println(err)
		}
		defer additionalTerritoriesXML.Close()

		additionalByteValue, _ := ioutil.ReadAll(additionalTerritoriesXML)

		var additionalTerritories Territories
		xml.Unmarshal(additionalByteValue, &additionalTerritories)

		territories.Territories = append(territories.Territories, additionalTerritories.Territories[0])
	}

	file, _ := xml.MarshalIndent(territories, "", "	")
	_ = ioutil.WriteFile("xml/zombie_territories.xml", file, 0644)
	color.Green("\n\nWrote `zombie_territories.xml` to the `xml/` directory. Upload to your server to modify spawns")
}
