# 🧟 ZombieMultiplierTool

Multiply the zombie count in the `zombie_territories.xml` files in DayZ SA automatically.

## 🗃️ Build Your Own File
1 - Update `xml/zombie_territories.xml` with the contents of your `zombie_territories.xml` file from your server. 
2 - Run **ZombieMultiplierTool** with the flags you want to use.
Example:
### ❕ `./ZombieMultiplierTool -InfectedGlobal=1.8`
3 - Your outputted file is at `xml/zombie_territories.xml`, replace the file on your server with this.
4 - Restart your DZSA server!

Additional multipliers can be added on top of the **InfectedGlobal** flag to increase ALL spawns by the global amount, then additional multipliers on top of that.

For example, if you'd like to increase the global spawn rate by 1.5, then increase the **InfectedArmy** spawn rates by another 2x multiplier, you'd run the command below, note this increases the **InfectedArmy** spawn rate to 2.5.

`./ZombieMultiplierTool -InfectedGlobal=1.5 -InfectedArmy=2`

### 🏳️ Flags
**AffectMin** - If set to true, minimum spawns will also be multiplied.
**Radius** - Set this to an amount to multiply the spawn radius by, this will spread out the spawns of zombies.
**Infected Types**...
``
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

## 🎯 Additional Spawns
If you'd like to add additional spawns, add your spawns to the `xml/additional_spawns.xml` file and build with the flag `-AdditionalSpawns=true`.
