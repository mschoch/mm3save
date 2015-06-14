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

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mschoch/mm3save"
)

var o = flag.String("o", "", "output filename")

var name = flag.String("name", "", "name of save game")
var food = flag.Int("food", -1, "party food")
var gold = flag.Int("gold", -1, "party gold")
var gems = flag.Int("gems", -1, "party gems")
var bankgold = flag.Int("bankgold", -1, "bank gold")
var bankgems = flag.Int("bankgems", -1, "bank gems")
var dir = flag.Int("dir", -1, "direction (0-N, 1-E, 2-S, 3-W)")
var x = flag.Int("x", -1, "x position")
var y = flag.Int("y", -2, "y position")

// character specific updates
var char = flag.String("char", "", "name of the character you want to update")

// character attributes
var mgt = flag.Int("mgt", -1, "might")
var intx = flag.Int("int", -1, "intellect")
var per = flag.Int("per", -1, "personality")
var end = flag.Int("end", -1, "endurance")
var spd = flag.Int("spd", -1, "speed")
var acy = flag.Int("acy", -1, "accuracy")
var lck = flag.Int("lck", -1, "luck")

// character resistances
var fire = flag.Int("fire", -1, "fire resistance")
var cold = flag.Int("cold", -1, "cold resistance")
var elec = flag.Int("elec", -1, "electric resistance")
var poison = flag.Int("poison", -1, "poison resistance")
var energy = flag.Int("energy", -1, "energey resistance")
var magic = flag.Int("magic", -1, "magic resistance")

var hp = flag.Int("hp", -1, "hit points")
var sp = flag.Int("sp", -1, "spell points")
var exp = flag.Int("exp", -1, "experience")

var skills = flag.String("skills", "", "skills")

var itemHelp = flag.Bool("itemHelp", false, "detailed help for items")
var itemSlot = flag.Int("itemSlot", -1, "item slot (1-18)")
var itemDesc = flag.String("itemDesc", "", "<element> <attribute> <substance> type <special> (run -itemHelp for details)")

func main() {

	flag.Parse()

	if *itemHelp {
		fmt.Printf("Items are built with strings of the form: <element> <attribute> <substance> type <special>\n")
		fmt.Printf("Only type is required, everything else is optional.\n\n")
		mm3save.ItemHelp()
		return
	}

	if flag.NArg() < 1 {
		fmt.Printf("filename required\n")
		return
	}

	if *o == "" {
		fmt.Printf("output filename is required\n")
		return
	}

	f, err := mm3save.NewFile(flag.Arg(0))
	if err != nil {
		fmt.Printf("error parsing file: %v\n", err)
		return
	}

	// open a temp file for output
	output, err := ioutil.TempFile("", "mm3update")
	if err != nil {
		fmt.Printf("error opening output file: %v\n", err)
		return
	}

	if *name != "" {
		if len(*name) > 24 {
			fmt.Printf("save game name cannot be more than 24 characters\n")
			return
		}
		fmt.Printf("setting save game name to '%s'\n", *name)
		f.Name = *name
	}

	if *food >= 0 {
		fmt.Printf("setting party food to %d\n", *food)
		f.Food = uint16(*food)
	}

	if *gold >= 0 {
		fmt.Printf("setting party gold to %d\n", *gold)
		f.Gold = uint32(*gold)
	}

	if *gems >= 0 {
		fmt.Printf("setting party gems to %d\n", *gems)
		f.Gems = uint32(*gems)
	}

	if *bankgold >= 0 {
		fmt.Printf("setting bank gold to %d\n", *bankgold)
		f.BankGold = uint32(*bankgold)
	}

	if *bankgems >= 0 {
		fmt.Printf("setting bank gems to %d\n", *bankgems)
		f.BankGems = uint32(*bankgems)
	}

	if *dir >= 0 {
		fmt.Printf("setting direction to %s (%d)\n", mm3save.Directions[byte(*dir)], *dir)
		f.Direction = byte(*dir)
	}

	if *x >= 0 {
		fmt.Printf("setting x position to %d\n", *x)
		f.X = byte(*x)
	}

	if *y >= 0 {
		fmt.Printf("setting y position to %d\n", *y)
		f.Y = byte(*y)
	}

	if *char != "" {
		found := false
		for _, character := range f.Characters {
			if character.Name == *char {
				found = true
				fmt.Printf("updating character '%s'\n", *char)
				if *mgt >= 0 {
					fmt.Printf("setting MGT to %d\n", *mgt)
					character.Mgt = byte(*mgt)
				}
				if *intx >= 0 {
					fmt.Printf("setting INT to %d\n", *intx)
					character.Int = byte(*intx)
				}
				if *per >= 0 {
					fmt.Printf("setting PER to %d\n", *per)
					character.Per = byte(*per)
				}
				if *end >= 0 {
					fmt.Printf("setting END to %d\n", *end)
					character.End = byte(*end)
				}
				if *spd >= 0 {
					fmt.Printf("setting SPD to %d\n", *spd)
					character.Spd = byte(*spd)
				}
				if *acy >= 0 {
					fmt.Printf("setting ACY to %d\n", *acy)
					character.Acy = byte(*acy)
				}
				if *lck >= 0 {
					fmt.Printf("setting LCK to %d\n", *lck)
					character.Lck = byte(*lck)
				}
				if *fire >= 0 {
					fmt.Printf("setting fire res to %d\n", *fire)
					character.ResFire = byte(*fire)
				}
				if *cold >= 0 {
					fmt.Printf("setting cold res to %d\n", *cold)
					character.ResCold = byte(*cold)
				}
				if *elec >= 0 {
					fmt.Printf("setting elec res to %d\n", *elec)
					character.ResElec = byte(*elec)
				}
				if *poison >= 0 {
					fmt.Printf("setting poison res to %d\n", *poison)
					character.ResPoison = byte(*poison)
				}
				if *energy >= 0 {
					fmt.Printf("setting energy res to %d\n", *energy)
					character.ResEnergy = byte(*energy)
				}
				if *magic >= 0 {
					fmt.Printf("setting magic res to %d\n", *magic)
					character.ResMagic = byte(*magic)
				}
				if *hp >= 0 {
					fmt.Printf("setting hp to %d\n", *hp)
					character.Hp = uint16(*hp)
				}
				if *sp >= 0 {
					fmt.Printf("setting sp to %d\n", *sp)
					character.Sp = uint16(*sp)
				}
				if *exp >= 0 {
					fmt.Printf("setting exp to %d\n", *exp)
					character.Exp = uint32(*exp)
				}
				if *skills != "" {
					fmt.Printf("giving character the skills %s\n", *skills)
					character.SetSkills(*skills)
				}
				if *itemSlot >= 1 && *itemSlot <= 18 {
					fmt.Printf("setting item %d to %s\n", *itemSlot, *itemDesc)
					item, err := mm3save.NewItem(*itemDesc)
					if err != nil {
						fmt.Printf("error parsing item %v\n", err)
						return
					}
					character.Items[*itemSlot-1] = item
				}
			}

		}
		if !found {
			fmt.Printf("could not find character '%s'\n", *char)
			return
		}
	}

	_, err = f.WriteTo(output)
	if err != nil {
		fmt.Printf("error writing output file: %v\n", err)
		return
	}

	// if we got this far, check the size of the output file
	fi, err := output.Stat()
	if err != nil {
		fmt.Printf("error getting stats of tmp output file: %v\n", err)
		return
	}

	if fi.Size() != 207551 {
		fmt.Printf("tmp output file size is invalid, %d not 207551, NOT renaming", fi.Size())
		return
	}

	// file size is OK, rename it to the requested output filename
	os.Rename(output.Name(), *o)

	err = output.Close()
	if err != nil {
		fmt.Printf("error closing output file: %v\n", err)
		return
	}
}
