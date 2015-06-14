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
	"fmt"
	"strings"
)

type Item struct {
	Slot      byte
	Substance byte
	Type      byte
	Elemental byte
	Attribute byte
	Special   byte
}

var UnknownItemErr = fmt.Errorf("unable to determine item type")

func NewItem(str string) (*Item, error) {
	rv := Item{}
	fields := strings.Fields(str)
	for _, f := range fields {

		lookup := strings.ToLower(f)
		lookup = strings.Replace(lookup, "_", " ", -1)

		if lookup == "of" {
			continue
		}
		sub, ok := invItemSubstances[lookup]
		if ok {
			rv.Substance = sub
			continue
		}
		typ, ok := invItemTypes[lookup]
		if ok {
			rv.Type = typ
			continue
		}
		elm, ok := invItemElementalProps[lookup]
		if ok {
			rv.Elemental = elm
			continue
		}
		attr, ok := invItemAttributeProps[lookup]
		if ok {
			rv.Attribute = attr
			continue
		}
		sp, ok := invItemSpecialProps["of "+lookup]
		if ok {
			rv.Special = sp
			continue
		}
	}
	if rv.Type == 0 {
		return nil, UnknownItemErr
	}
	return &rv, nil
}

func (i *Item) String() string {
	rv := ""
	if i.Elemental > 0 {
		rv += itemElementalProps[i.Elemental]
	}
	if i.Substance > 0 {
		if len(rv) > 0 {
			rv += " "
		}
		rv += itemSubstances[i.Substance]
	}
	if i.Attribute > 0 {
		if len(rv) > 0 {
			rv += " "
		}
		rv += itemAttributeProps[i.Attribute]
	}
	if i.Type > 0 {
		if len(rv) > 0 {
			rv += " "
		}
		rv += itemTypes[i.Type]
	}
	if i.Special > 0 {
		if len(rv) > 0 {
			rv += " "
		}
		rv += itemSpecialProps[i.Special]
	}
	if i.Slot > 0 {
		if len(rv) > 0 {
			rv += " "
		}
		rv += "(equipped as " + itemEquippedSLot[i.Slot] + ")"
	}
	return rv
}

const unequipped = 0
const equippedLeftHand = 1
const equippedRightHand = 2
const equippedArmor = 3
const equippedRangeWeapon = 4
const equippedHeadgear = 5
const equippedGuantlets = 6
const equippedMedal = 7
const equippedRing = 8
const equippedBoot = 9
const equippedCloak = 10
const equipedNecklace = 11
const equippedBelt = 12
const equippedTwoHanded = 13

var itemEquippedSLot = map[byte]string{
	1:  "Left Hand",
	2:  "Right Hand",
	3:  "Armor",
	4:  "RangedWeapon",
	5:  "Headgear",
	6:  "Guantlets",
	7:  "Medal",
	8:  "Ring",
	9:  "Boots",
	10: "Cloak",
	11: "Necklace",
	12: "Belt",
	13: "Two Handed",
}

var invItemTypes = map[string]byte{}

var itemTypes = map[byte]string{
	1:   "Long Sword",
	2:   "Short Sword",
	3:   "Broad Sword",
	4:   "Scimitar",
	5:   "Cutlass",
	6:   "Sabre",
	7:   "Club",
	8:   "Hand Axe",
	9:   "Katana",
	10:  "Nunchakas",
	11:  "Wakazashi",
	12:  "Dagger",
	13:  "Mace",
	14:  "Flail",
	15:  "Cudgel",
	16:  "Maul",
	17:  "Spear",
	18:  "Bardiche",
	19:  "Glaive",
	20:  "Halberd",
	21:  "Pike",
	22:  "Flamberge",
	23:  "Trident",
	24:  "Staff",
	25:  "Hammer",
	26:  "Naginata",
	27:  "Battle Axe",
	28:  "Grand Axe",
	29:  "Great Axe",
	30:  "Short Bow",
	31:  "Long Bow",
	32:  "Crossbow",
	33:  "Sling",
	34:  "Padded Armor",
	35:  "Leather Armor",
	36:  "Scale Armor",
	37:  "Ring Mail",
	38:  "Chain Mail",
	39:  "Splint Mail",
	40:  "Plate Mail",
	41:  "Plate Armor",
	42:  "Shield",
	43:  "Helm",
	44:  "Crown",
	45:  "Tiara",
	46:  "Gauntlets",
	47:  "Ring",
	48:  "Boots",
	49:  "Cloak",
	50:  "Robes",
	51:  "Cape",
	52:  "Belt",
	53:  "Broach",
	54:  "Medal",
	55:  "Charm",
	56:  "Cameo",
	57:  "Scarab",
	58:  "Pendant",
	59:  "Necklace",
	60:  "Amulet",
	61:  "Rod",
	62:  "Jewel",
	63:  "Gem",
	64:  "Box",
	65:  "Orb",
	66:  "Horn",
	67:  "Coin",
	68:  "Wand",
	69:  "Whistle",
	70:  "Potion",
	71:  "Scroll",
	72:  "Torch",
	73:  "Rope and Hooks",
	74:  "Useless Item",
	75:  "Jewelry, Ancient",
	76:  "Green Eyeball Key",
	77:  "Red Warriors Key",
	78:  "Sacred Silver Skull",
	79:  "Ancient Artifact of Good",
	80:  "Ancient Artifact of Neutrality",
	81:  "Ancient Artifact of Evil",
	82:  "Jewelry",
	83:  "Precious Pearl of Youth and Beauty",
	84:  "Black Terror Key",
	85:  "King's Ultimate Power Orb",
	86:  "Ancient Fizbin of Misfortune",
	87:  "Gold Master Key",
	88:  "Quatloo Coin",
	89:  "Hologram Sequencing Card 001",
	90:  "Yellow Fortress Key",
	91:  "Blue Unholy Key",
	92:  "Hologram Sequencing Card 002",
	93:  "Hologram Sequencing Card 003",
	94:  "Hologram Sequencing Card 004",
	95:  "Hologram Sequencing Card 005",
	96:  "Hologram Sequencing Card 006",
	97:  "Z Item 23",
	98:  "Blue Priority Pass Card",
	99:  "Interspatial Transport Box",
	100: "Might Potion",
	101: "Golden Pyramid Access Card",
	102: "Alacorn of Icarus",
	103: "Sea Shell of Serenity",
}

var invItemSubstances = map[string]byte{}

var itemSubstances = map[byte]string{
	1:  "Wooden",
	2:  "Leather",
	3:  "Brass",
	4:  "Bronze",
	5:  "Iron",
	6:  "Silver",
	7:  "Steel",
	8:  "Gold",
	9:  "Platinum",
	10: "Glass",
	11: "Coral",
	12: "Crystal",
	13: "Lapis",
	14: "Pearl",
	15: "Amber",
	16: "Ebony",
	17: "Quartz",
	18: "Ruby",
	19: "Emerald",
	20: "Sapphire",
	21: "Diamond",
	22: "Obsidian",
}

var invItemElementalProps = map[string]byte{}

var itemElementalProps = map[byte]string{
	1:  "Burning",
	2:  "Fiery",
	3:  "Pyric",
	4:  "Fuming",
	5:  "Flaming",
	6:  "Seething",
	7:  "Blazing",
	8:  "Scorching",
	9:  "Flickering",
	10: "Sparking",
	11: "Static",
	12: "Flashing",
	13: "Shocking",
	14: "Electric",
	15: "Dyna",
	16: "Icy",
	17: "Frost",
	18: "Freezing",
	19: "Cold",
	20: "Cryo",
	21: "Acidic",
	22: "Venemous",
	23: "Poisonous",
	24: "Toxic",
	25: "Noxious",
	26: "Glowing",
	27: "Incandescent",
	28: "Dense",
	29: "Sonic",
	30: "Power",
	31: "Thermal",
	32: "Radiating",
	33: "Kinetic",
	34: "Mystic",
	35: "Magical",
	36: "Ectoplasmic",
}

var invItemAttributeProps = map[string]byte{}

var itemAttributeProps = map[byte]string{
	1:  "Might",
	2:  "Strength",
	3:  "Warrior",
	4:  "Ogre",
	5:  "Giant",
	6:  "Thunder",
	7:  "Force",
	8:  "Power",
	9:  "Dragon",
	10: "Photon",
	11: "Clever",
	12: "Mind",
	13: "Sage",
	14: "Thought",
	15: "Knowledge",
	16: "Intellect",
	17: "Wisdom",
	18: "Genius",
	19: "Buddy",
	20: "Friendship",
	21: "Charm",
	22: "Personality",
	23: "Charisma",
	24: "Leadership",
	25: "Ego",
	26: "Holy",
	27: "Quick",
	28: "Swift",
	29: "Fast",
	30: "Rapid",
	31: "Speed",
	32: "Wind",
	33: "Accelerator",
	34: "Velocity",
	35: "Sharp",
	36: "Accurate",
	37: "Marksman",
	38: "Precision",
	39: "True",
	40: "Exacto",
	41: "Clover",
	42: "Chance",
	43: "Winners",
	44: "Lucky",
	45: "Gamblers",
	46: "Leprachauns",
	47: "Vigor",
	48: "Health",
	49: "Life",
	50: "Troll",
	51: "Vampiric",
	52: "Spell",
	53: "Casters",
	54: "Witch",
	55: "Mage",
	56: "Archmage",
	57: "Arcane",
	58: "Protection",
	59: "Armored",
	60: "Defender",
	61: "Stealth",
	62: "Divine",
	63: "Mugger",
	64: "Burgler",
	65: "Looter",
	66: "Brigand",
	67: "Filch",
	68: "Thief",
	69: "Rogue",
	70: "Plunder",
	71: "Criminal",
	72: "Pirate",
}

var invItemSpecialProps = map[string]byte{}

var itemSpecialProps = map[byte]string{
	1:  "of Light",
	2:  "of Awakening",
	3:  "of Magic Detection",
	4:  "of Arrows",
	5:  "of Aid",
	6:  "of Fists",
	7:  "of Energy Blasts",
	8:  "of Sleeping",
	9:  "of Revitalization",
	10: "of Curing",
	11: "of Sparking",
	12: "of Ropes",
	13: "of Toxic Clouds",
	14: "of Elements",
	15: "of Pain",
	16: "of Jumping",
	17: "of Acid Stream",
	18: "of Undead Turning",
	19: "of Levitation",
	20: "of Wizard Eyes",
	21: "of Silence",
	22: "of Blessing",
	23: "of Identification",
	24: "of Lightning",
	25: "of Holy Bonuses",
	26: "of Power Cures",
	27: "of Nature",
	28: "of Beacons",
	29: "of Shielding",
	30: "of Heroism",
	31: "of Immobilization",
	32: "of Water Walking",
	33: "of Frost Biting",
	34: "of Monster Finding",
	35: "of Fireballs",
	36: "of Cold Rays",
	37: "of Antidotes",
	38: "of Acid Spraying",
	39: "of Distortion",
	40: "of Feeble Minding",
	41: "of Vaccination",
	42: "of Gating",
	43: "of Teleportation",
	44: "of Death",
	45: "of Free Movement",
	46: "of Paralyzing",
	47: "of Deadly Swarms",
	48: "of Sanctuaries",
	49: "of Dragon Breath",
	50: "of Feasting",
	51: "of Fiery Flails",
	52: "of Recharging",
	53: "of Freezing",
	54: "of Portals",
	55: "of Stone to Flesh",
	56: "of Duplication",
	57: "of Disintegration",
	58: "of Half for Me",
	59: "of Raising the Dead",
	60: "of Etherealization",
	61: "of Dancing Swords",
	62: "of Moon Rays",
	63: "of Mass Distortion",
	64: "of Prismatic Light",
	65: "of Enchantment",
	66: "of Incinerating",
	67: "of Holy Words",
	68: "of Resurrection",
	69: "of Storms",
	70: "of Megavoltage",
	71: "of Infernos",
	72: "of Sun Rays",
	73: "of Implosion",
	74: "of Star Bursts",
	75: "of the GODS!",
}

func init() {
	// invert lookup tables
	for k, v := range itemTypes {
		invItemTypes[strings.ToLower(v)] = k
	}
	for k, v := range itemSubstances {
		invItemSubstances[strings.ToLower(v)] = k
	}
	for k, v := range itemElementalProps {
		invItemElementalProps[strings.ToLower(v)] = k
	}
	for k, v := range itemAttributeProps {
		invItemAttributeProps[strings.ToLower(v)] = k
	}
	for k, v := range itemSpecialProps {
		invItemSpecialProps[strings.ToLower(v)] = k
	}
}

func ItemHelp() {
	fmt.Printf("Elements:\n\n")
	for _, v := range itemElementalProps {
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("\nSubstances:\n\n")
	for _, v := range itemSubstances {
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("\nAttributes:\n\n")
	for _, v := range itemAttributeProps {
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("\nTypes:\n\n")
	for _, v := range itemTypes {
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("\nSpecials:\n\n")
	for _, v := range itemSpecialProps {
		fmt.Printf("%s\n", v)
	}
}
