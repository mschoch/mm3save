//  Copyright (c) 2015 Marty Schoch
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the
//  License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an "AS
//  IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
//  express or implied. See the License for the specific language
//  governing permissions and limitations under the License.

package mm3save

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

const firstCharPos = 1953
const numCharacters = 30
const maxCharName = 10

const genderOffset = 16
const raceOffet = 17
const alignmentOffset = 18
const classOffset = 19

const mgtOffset = 20
const intOffset = mgtOffset + 2
const perOffset = intOffset + 2
const endOffset = perOffset + 2
const spdOffset = endOffset + 2
const acyOffset = spdOffset + 2
const lckOffset = acyOffset + 2

const lvlOffset = 35

const skillThieveryOffset = 39 //?
const skillArmsMasterOffset = 40
const skillAstrologerOffset = 41
const skillBodyBuilderOffset = 42
const skillCartographyOffset = 43
const skillCrusaderOffset = 44
const skillDirectionSenseOffset = 45
const skillLinguistOffst = 46
const skillMerchantOffset = 47
const skillMountaineerOffset = 48
const skillNavigatorOffset = 49
const skillPathfinderOffset = 50
const skillPrayerMasterOffset = 51
const skillPrestidigitatorOffset = 52
const skillSwimmerOffset = 53
const skillTrackerOffset = 54
const skillSpotSecretDoorsOffset = 55
const skillDangerSenseOffset = 56

const awardRavensGuildMember = 57
const awardAlbatrossGuildMember = 58
const awardFalconGuildMember = 59
const awardBuzzardGuildMember = 60
const awardEagleGuildMember = 61
const awardSavedFountainHead = 62
const awaredBlessedByTheForces = 63
const awardNumOrbsGivenToZealout = 64
const awardNumOrbsGivenToMalefactor = 65
const awardNumOrbsGivenToTumult = 66
const awardUnknown = 67
const awardNumGoodArtifactsRecovered = 68
const awardNumEvilArtifactsRecovered = 69
const awardNumNeutralArtifactsRecovered = 70
const awardNumShellsGivenToAthena = 71
const awardGreekBrothersVisited = 72
const awardGreywindReleased645 = 73
const awardBlackwindReleased231 = 74
const awardNumSkullGivenToKranion = 75
const awardIcarusResurected = 76
const awardFreedPrincessTrueberry = 77
const awardNumArenaWins = 78
const awardNumPearlsToPirateQueen = 79
const awardUltimateAdventurer = 80

const firstItemEquipped = 125
const firstItemUnknown = firstItemEquipped + 19
const firstItemElemental = firstItemUnknown + 19
const firstItemSubtance = firstItemElemental + 19
const firstItemAttribute = firstItemSubtance + 19
const firstItemType = firstItemAttribute + 19
const firstItemSpecial = firstItemType + 19

const fireResOffset = 263
const coldResOffset = fireResOffset + 2
const elecResOffset = coldResOffset + 2
const poisonResOffset = elecResOffset + 2
const energyResOffset = poisonResOffset + 2
const magicResOffset = energyResOffset + 2

const hpOffset = 293
const spOffset = 295

const expOffset = 299
const charSize = 303

func ParseCharacters(in []byte) ([]*Char, error) {
	rv := make([]*Char, numCharacters)
	pos := firstCharPos
	for i := 0; i < numCharacters; i++ {
		charData := in[pos : pos+charSize]
		char, err := NewChar(charData)
		if err != nil {
			return nil, err
		}
		rv[i] = char
		pos += charSize
	}
	return rv, nil
}

type Char struct {
	Orig []byte
	Name string

	Gender    byte
	Race      byte
	Alignment byte
	Class     byte

	Lvl byte

	Exp uint32

	Mgt    byte
	MgtMod byte
	Int    byte
	IntMod byte
	Per    byte
	PerMod byte
	End    byte
	EndMod byte
	Spd    byte
	SpdMod byte
	Acy    byte
	AcyMod byte
	Lck    byte
	LckMod byte

	Age uint16
	Hp  uint16
	Sp  uint16

	// res
	ResFire      byte
	ResFireMod   byte
	ResCold      byte
	ResColdMod   byte
	ResElec      byte
	ResElecMod   byte
	ResPoison    byte
	ResPoisonMod byte
	ResEnergy    byte
	ResEnergyMod byte
	ResMagic     byte
	ResMagicMod  byte

	SkillThievery        byte
	SkillArmsMaster      byte
	SkillAstrologer      byte
	SkillBodyBuilder     byte
	SkillCartography     byte
	SkillCrusader        byte
	SkillDirectionSense  byte
	SkillLinguist        byte
	SkillMerchant        byte
	SkillMountaineer     byte
	SkillNavigator       byte
	SkillPathfinder      byte
	SkillPrayerMaster    byte
	SkillPrestidigitator byte
	SkillSwimmer         byte
	SkillTracker         byte
	SkillSpotSecretDoors byte
	SkillDangerSense     byte

	Items []*Item
}

func (c *Char) WriteTo(w io.Writer) (n int64, err error) {
	// name at most 10 bytes, zero'd out not orig (since you could have shortened it)
	written := int64(0)
	nameBytes := make([]byte, maxCharName)
	copy(nameBytes, []byte(c.Name))
	if len(c.Name) < maxCharName {
		nameBytes[len(c.Name)] = 0
	}
	wrote, err := w.Write(nameBytes)
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write Orig after name
	wrote, err = w.Write(c.Orig[maxCharName:genderOffset])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write core
	wrote, err = w.Write([]byte{c.Gender, c.Race, c.Alignment, c.Class})
	written += int64(wrote)
	if err != nil {
		return written, err
	}
	// write attributes
	wrote, err = w.Write([]byte{
		c.Mgt,
		c.MgtMod,
		c.Int,
		c.IntMod,
		c.Per,
		c.PerMod,
		c.End,
		c.EndMod,
		c.Spd,
		c.SpdMod,
		c.Acy,
		c.AcyMod,
		c.Lck,
		c.LckMod,
	})
	written += int64(wrote)
	if err != nil {
		return written, err
	}
	// write gap before level
	wrote, err = w.Write(c.Orig[lckOffset+2 : lckOffset+3])
	written += int64(wrote)
	if err != nil {
		return written, err
	}
	// write level
	wrote, err = w.Write([]byte{c.Lvl})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out the rest up til first skill
	wrote, err = w.Write(c.Orig[lvlOffset+1 : skillThieveryOffset])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out the skills
	wrote, err = w.Write([]byte{
		c.SkillThievery,
		c.SkillArmsMaster,
		c.SkillAstrologer,
		c.SkillBodyBuilder,
		c.SkillCartography,
		c.SkillCrusader,
		c.SkillDirectionSense,
		c.SkillLinguist,
		c.SkillMerchant,
		c.SkillMountaineer,
		c.SkillNavigator,
		c.SkillPathfinder,
		c.SkillPrayerMaster,
		c.SkillPrestidigitator,
		c.SkillSwimmer,
		c.SkillTracker,
		c.SkillSpotSecretDoors,
		c.SkillDangerSense,
	})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out the rest up til items
	wrote, err = w.Write(c.Orig[skillDangerSenseOffset+1 : firstItemEquipped])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out item equipped status
	for _, item := range c.Items {
		wrote, err = w.Write([]byte{item.Slot})
		written += int64(wrote)
		if err != nil {
			return written, err
		}
	}
	// write out 0 termination
	wrote, err = w.Write([]byte{0})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out item unknown
	for range c.Items {
		wrote, err = w.Write([]byte{0})
		written += int64(wrote)
		if err != nil {
			return written, err
		}
	}
	// write out 0 termination
	wrote, err = w.Write([]byte{0})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out item element
	for _, item := range c.Items {
		wrote, err = w.Write([]byte{item.Elemental})
		written += int64(wrote)
		if err != nil {
			return written, err
		}
	}
	// write out 0 termination
	wrote, err = w.Write([]byte{0})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out item substance
	for _, item := range c.Items {
		wrote, err = w.Write([]byte{item.Substance})
		written += int64(wrote)
		if err != nil {
			return written, err
		}
	}
	// write out 0 termination
	wrote, err = w.Write([]byte{0})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out item attribute
	for _, item := range c.Items {
		wrote, err = w.Write([]byte{item.Attribute})
		written += int64(wrote)
		if err != nil {
			return written, err
		}
	}
	// write out 0 termination
	wrote, err = w.Write([]byte{0})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out item type
	for _, item := range c.Items {
		wrote, err = w.Write([]byte{item.Type})
		written += int64(wrote)
		if err != nil {
			return written, err
		}
	}
	// write out 0 termination
	wrote, err = w.Write([]byte{0})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out item special
	for _, item := range c.Items {
		wrote, err = w.Write([]byte{item.Special})
		written += int64(wrote)
		if err != nil {
			return written, err
		}
	}
	// write out 0 termination
	wrote, err = w.Write([]byte{0})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write out the rest up til resistences
	wrote, err = w.Write(c.Orig[firstItemSpecial+19 : fireResOffset])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// now write out the resistances
	wrote, err = w.Write([]byte{
		c.ResFire,
		c.ResFireMod,
		c.ResCold,
		c.ResColdMod,
		c.ResElec,
		c.ResElecMod,
		c.ResPoison,
		c.ResPoisonMod,
		c.ResEnergy,
		c.ResEnergyMod,
		c.ResMagic,
		c.ResMagicMod,
	})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// more from the original
	wrote, err = w.Write(c.Orig[magicResOffset+2 : hpOffset])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write hp
	hpBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(hpBytes, c.Hp)
	wrote, err = w.Write(hpBytes)
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write sp
	spBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(spBytes, c.Sp)
	wrote, err = w.Write(spBytes)
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// more from the original
	wrote, err = w.Write(c.Orig[spOffset+2 : expOffset])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write exp
	expBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(expBytes, c.Exp)
	wrote, err = w.Write(expBytes)
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	return written, nil
}

func (c *Char) SetSkills(skills string) {
	sa := strings.Split(skills, ",")
	c.SkillThievery = 0
	c.SkillArmsMaster = 0
	c.SkillAstrologer = 0
	c.SkillBodyBuilder = 0
	c.SkillCartography = 0
	c.SkillCrusader = 0
	c.SkillDirectionSense = 0
	c.SkillLinguist = 0
	c.SkillMerchant = 0
	c.SkillMountaineer = 0
	c.SkillNavigator = 0
	c.SkillPathfinder = 0
	c.SkillPrayerMaster = 0
	c.SkillPrestidigitator = 0
	c.SkillSwimmer = 0
	c.SkillTracker = 0
	c.SkillSpotSecretDoors = 0
	c.SkillDangerSense = 0
	for _, s := range sa {
		if s == "thievery" {
			c.SkillThievery = 1
		}
		if s == "armsmaster" {
			c.SkillArmsMaster = 1
		}
		if s == "astrologer" {
			c.SkillAstrologer = 1
		}
		if s == "bodybuilder" {
			c.SkillBodyBuilder = 1
		}
		if s == "cartographer" {
			c.SkillCartography = 1
		}
		if s == "crusader" {
			c.SkillCrusader = 1
		}
		if s == "directionsense" {
			c.SkillDirectionSense = 1
		}
		if s == "linguist" {
			c.SkillLinguist = 1
		}
		if s == "merchant" {
			c.SkillMerchant = 1
		}
		if s == "mountaineer" {
			c.SkillMountaineer = 1
		}
		if s == "navigator" {
			c.SkillNavigator = 1
		}
		if s == "pathfinder" {
			c.SkillPathfinder = 1
		}
		if s == "prayermaster" {
			c.SkillPrayerMaster = 1
		}
		if s == "prestidigitator" {
			c.SkillPrestidigitator = 1
		}
		if s == "swimmer" {
			c.SkillSwimmer = 1
		}
		if s == "tracker" {
			c.SkillTracker = 1
		}
		if s == "spotsecretdoors" {
			c.SkillSpotSecretDoors = 1
		}
		if s == "dangersense" {
			c.SkillDangerSense = 1
		}
	}
}

func (c *Char) SkillsString() string {
	rv := []string{}
	if c.SkillThievery > 0 {
		rv = append(rv, "thievery")
	}
	if c.SkillArmsMaster > 0 {
		rv = append(rv, "armsmaster")
	}
	if c.SkillAstrologer > 0 {
		rv = append(rv, "astrologer")
	}
	if c.SkillBodyBuilder > 0 {
		rv = append(rv, "bodybuilder")
	}
	if c.SkillCartography > 0 {
		rv = append(rv, "cartographer")
	}
	if c.SkillCrusader > 0 {
		rv = append(rv, "crusader")
	}
	if c.SkillDirectionSense > 0 {
		rv = append(rv, "directionsense")
	}
	if c.SkillLinguist > 0 {
		rv = append(rv, "linguist")
	}
	if c.SkillMerchant > 0 {
		rv = append(rv, "merchant")
	}
	if c.SkillMountaineer > 0 {
		rv = append(rv, "mountaineer")
	}
	if c.SkillNavigator > 0 {
		rv = append(rv, "navigator")
	}
	if c.SkillPathfinder > 0 {
		rv = append(rv, "pathfinder")
	}
	if c.SkillPrayerMaster > 0 {
		rv = append(rv, "prayermaster")
	}
	if c.SkillPrestidigitator > 0 {
		rv = append(rv, "prestidigitator")
	}
	if c.SkillSwimmer > 0 {
		rv = append(rv, "swimmer")
	}
	if c.SkillTracker > 0 {
		rv = append(rv, "tracker")
	}
	if c.SkillSpotSecretDoors > 0 {
		rv = append(rv, "spotsecretdoors")
	}
	if c.SkillDangerSense > 0 {
		rv = append(rv, "dangersense")
	}
	return strings.Join(rv, ",")
}

func (c *Char) String() string {
	rv := ""
	if c.Name != "" {
		rv += fmt.Sprintf("Name: %s\n", c.Name)
		if c.Gender > 0 {
			rv += fmt.Sprintf("Gender: Female\n")
		} else {
			rv += fmt.Sprintf("Gender: Male\n")
		}
		rv += fmt.Sprintf("Race: %s\n", races[c.Race])
		rv += fmt.Sprintf("Alignment: %s\n", alignments[c.Alignment])
		rv += fmt.Sprintf("Class: %s\n", classes[c.Class])
		rv += fmt.Sprintf("Mgt: %d (+%d)\n", c.Mgt, c.MgtMod)
		rv += fmt.Sprintf("Int: %d (+%d)\n", c.Int, c.IntMod)
		rv += fmt.Sprintf("Per: %d (+%d)\n", c.Per, c.PerMod)
		rv += fmt.Sprintf("End: %d (+%d)\n", c.End, c.EndMod)
		rv += fmt.Sprintf("Spd: %d (+%d)\n", c.Spd, c.SpdMod)
		rv += fmt.Sprintf("Acy: %d (+%d)\n", c.Acy, c.AcyMod)
		rv += fmt.Sprintf("Lck: %d (+%d)\n", c.Lck, c.LckMod)
		rv += fmt.Sprintf("Exp: %d\n", c.Exp)
		rv += fmt.Sprintf("Lvl: %d\n", c.Lvl)

		rv += fmt.Sprintf("Res Fire: %d (+%d)\n", c.ResFire, c.ResFireMod)
		rv += fmt.Sprintf("Res Cold: %d (+%d)\n", c.ResCold, c.ResColdMod)
		rv += fmt.Sprintf("Res Elec: %d (+%d)\n", c.ResElec, c.ResElecMod)
		rv += fmt.Sprintf("Res Poison: %d (+%d)\n", c.ResPoison, c.ResPoisonMod)
		rv += fmt.Sprintf("Res Energy: %d (+%d)\n", c.ResEnergy, c.ResEnergyMod)
		rv += fmt.Sprintf("Res Magic: %d (+%d)\n", c.ResMagic, c.ResMagicMod)

		rv += fmt.Sprintf("HP: %d\n", c.Hp)
		rv += fmt.Sprintf("SP: %d\n", c.Sp)

		rv += fmt.Sprintf("Skills: %s\n", c.SkillsString())

		if len(c.Items) > 0 {
			rv += fmt.Sprintf("Items:\n")
			for i, item := range c.Items {
				if item != nil {
					rv += fmt.Sprintf("%d - %s\n", i, item)
				}
			}
		}
	} else {
		rv += "Undefined Character\n"
	}
	return rv
}

func NewChar(in []byte) (*Char, error) {
	rv := &Char{
		Orig:  in,
		Items: make([]*Item, 18),
	}
	for i := 0; i < 18; i++ {
		rv.Items[i] = &Item{}
	}
	nameNullTerm := bytes.IndexByte(in, 0)
	if nameNullTerm > 0 {
		rv.Name = string(in[0:nameNullTerm])

		rv.Gender = in[genderOffset]
		rv.Race = in[raceOffet]
		rv.Alignment = in[alignmentOffset]
		rv.Class = in[classOffset]

		rv.Mgt = in[mgtOffset]
		rv.MgtMod = in[mgtOffset+1]
		rv.Int = in[intOffset]
		rv.IntMod = in[intOffset+1]
		rv.Per = in[perOffset]
		rv.PerMod = in[perOffset+1]
		rv.End = in[endOffset]
		rv.EndMod = in[endOffset+1]
		rv.Spd = in[spdOffset]
		rv.SpdMod = in[spdOffset+1]
		rv.Acy = in[acyOffset]
		rv.AcyMod = in[acyOffset+1]
		rv.Lck = in[lckOffset]
		rv.LckMod = in[lckOffset+1]

		rv.Lvl = in[lvlOffset]

		// skillz
		rv.SkillThievery = in[skillThieveryOffset]
		rv.SkillArmsMaster = in[skillArmsMasterOffset]
		rv.SkillAstrologer = in[skillAstrologerOffset]
		rv.SkillBodyBuilder = in[skillBodyBuilderOffset]
		rv.SkillCartography = in[skillCartographyOffset]
		rv.SkillCrusader = in[skillCrusaderOffset]
		rv.SkillDirectionSense = in[skillDirectionSenseOffset]
		rv.SkillLinguist = in[skillLinguistOffst]
		rv.SkillMerchant = in[skillMerchantOffset]
		rv.SkillMountaineer = in[skillMountaineerOffset]
		rv.SkillNavigator = in[skillNavigatorOffset]
		rv.SkillPathfinder = in[skillPathfinderOffset]
		rv.SkillPrestidigitator = in[skillPrestidigitatorOffset]
		rv.SkillPrayerMaster = in[skillPrayerMasterOffset]
		rv.SkillSpotSecretDoors = in[skillSpotSecretDoorsOffset]
		rv.SkillSwimmer = in[skillSwimmerOffset]
		rv.SkillTracker = in[skillTrackerOffset]
		rv.SkillDangerSense = in[skillDangerSenseOffset]

		rv.ResFire = in[fireResOffset]
		rv.ResFireMod = in[fireResOffset+1]
		rv.ResCold = in[coldResOffset]
		rv.ResColdMod = in[coldResOffset+1]
		rv.ResElec = in[elecResOffset]
		rv.ResElecMod = in[elecResOffset+1]
		rv.ResPoison = in[poisonResOffset]
		rv.ResPoisonMod = in[poisonResOffset+1]
		rv.ResEnergy = in[energyResOffset]
		rv.ResEnergyMod = in[energyResOffset+1]
		rv.ResMagic = in[magicResOffset]
		rv.ResMagicMod = in[magicResOffset+1]

		// lets use the type as primary
		// if type is 0, ignore other fields
		for i := 0; i < 18; i++ {
			t := in[firstItemType+i]
			if t > 0 {
				// build an item
				item := &Item{
					Slot:      in[firstItemEquipped+i],
					Elemental: in[firstItemElemental+i],
					Substance: in[firstItemSubtance+i],
					Attribute: in[firstItemAttribute+i],
					Type:      t,
					Special:   in[firstItemSpecial+i],
				}
				rv.Items[i] = item
			}
		}

		rv.Hp = binary.LittleEndian.Uint16(in[hpOffset : hpOffset+2])
		rv.Sp = binary.LittleEndian.Uint16(in[spOffset : spOffset+2])

		rv.Exp = binary.LittleEndian.Uint32(in[expOffset : expOffset+4])

	}
	return rv, nil
}

var races = map[byte]string{
	0: "Human",
	1: "Elf",
	2: "Gnome",
	3: "Dwarf",
	4: "Half-Orc",
}

var alignments = map[byte]string{
	0: "Good",
	1: "Neutral",
	2: "Evil",
}

var classes = map[byte]string{
	0: "Knight",
	1: "Paladin",
	2: "Archer",
	3: "Cleric",
	4: "Sorcerer",
	5: "Robber",
	6: "Ninja",
	7: "Barbarian",
	8: "Druid",
	9: "Ranger",
}
