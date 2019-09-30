package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"git.r.etro.sh/RetroPronghorn/ZombieMultiplierTool/src/zone"
	"github.com/fatih/color"
)

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

	var territories zone.Territories
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
			// Apply multipliers to zone
			selectedZones[j].MultiplyZone(
				*flags.InfectedArmy+*flags.InfectedGlobal-1,
				*flags.Radius,
				*flags.AffectMin,
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

		var additionalTerritories zone.Territories
		xml.Unmarshal(additionalByteValue, &additionalTerritories)

		territories.Territories = append(territories.Territories, additionalTerritories.Territories[0])
	}

	// Write modified file
	file, _ := xml.MarshalIndent(territories, "", "	")
	_ = ioutil.WriteFile("xml/zombie_territories.xml", file, 0644)
	// Done
	color.Green("\n\nWrote `zombie_territories.xml` to the `xml/` directory. Upload to your server to modify spawns")
}
