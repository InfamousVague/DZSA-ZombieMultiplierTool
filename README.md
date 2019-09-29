# ZombieMultiplierTool

Multiply the zombie count in the `zombie_territories.xml` files in DayZ SA programatically.

## Build Your Own File
To build your own file, you'll need to run this application with flags corrispoinding to your multiplier you'd like to apply. For example to increase the spawn amount of `InfectedArmy` zombies to 2x default, you'd add the flag `-InfectedArmy=2`.

`./ZombieMultiplierTool -InfectedGlobal=2 -InfectedArmy=3 -InfectedPolice=2`

Flag Types:
```
InfectedArmy
InfectedVillage
InfectedMedic
InfectedPolice
InfectedReligious
InfectedIndustrial
InfectedFirefighter
InfectedCity
InfectedSolitude
```

You can also use the global multiplier flag to set a base multiplier (if you just want to 2x every type of zombie, you can set this flag and ignore the rest). For example `-InfectedGlobal=2`.

## Input and Output files
Set the `xml/zombie_territories.base.xml` to the contents of your current zombie_territories file from your server. The output file will be `xml/zombie_territories.xml` located after the program runs and modifies your xml base.