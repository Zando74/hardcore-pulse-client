package entity

import (
	"strconv"
	"strings"
)

type DeathRecord struct {
	ClassID   	int     
	Guild     	string  
	Date      	int64   
	MapPos    	string  
	SourceID  	int     
	Name      	string  
	LastWords 	string  
	Level     	int     
	MapID     	int
	InstanceID 	int
	RaceID    	int   
	Realm 	  	string
}

var ClassNames = map[int]string{
	1: "WARRIOR",
	2: "PALADIN",
	3: "HUNTER",
	4: "ROGUE",
	5: "PRIEST",
	7: "SHAMAN",
	8: "MAGE",
	9: "WARLOCK",
	11: "DRUID",
}

var RaceNames = map[int]string{
	1: "Human",
	2: "Orc",
	3: "Dwarf",
	4: "Night Elf",
	5: "Undead",
	6: "Tauren",
	7: "Gnome",
	8: "Troll",
	10: "Blood Elf",
	11: "Draenei",
}

func (d *DeathRecord) GetXPos() float32 {
	if len(d.MapPos) == 0 {
		return 0
	}
	xPos, _ := strconv.ParseFloat(strings.Split(d.MapPos, ",")[0], 32)
	return float32(xPos)
}

func (d *DeathRecord) GetYPos() float32 {
	if len(d.MapPos) == 0 {
		return 0
	}
	yPos, _ := strconv.ParseFloat(strings.Split(d.MapPos, ",")[1], 32)
	return float32(yPos)
}

func (d *DeathRecord) GetClassName() string {
	return ClassNames[d.ClassID]
}

func (d *DeathRecord) GetRaceName() string {
	return RaceNames[d.RaceID]
}

func (d *DeathRecord) GetFactionName() string {
	if d.RaceID == 2 || d.RaceID == 5 || d.RaceID == 6 || d.RaceID == 8 || d.RaceID == 10 {
		return "Horde"
	}
	return "Alliance"
}

func (d *DeathRecord) GetMapID() int {
	if d.InstanceID != 0 {
		return d.InstanceID
	}
	return d.MapID
}