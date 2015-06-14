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
	"io/ioutil"
)

type File struct {
	Name       string
	Orig       []byte
	Direction  byte
	X          byte
	Y          byte
	BankGold   uint32
	BankGems   uint32
	Gold       uint32
	Gems       uint32
	Food       uint16
	Characters []*Char
}

const fileNameOffset = 1922
const maxFileName = 24
const partyDirectionOffset = 11053
const partyXPos = 11054
const partyYPos = 11055
const partyFoodOffset = 11901
const bankGoldOffset = 11909
const bankGemOffset = 11913
const partyGoldOffset = 11917
const partyGemOffset = 11921
const correctFileSize = 207551

func (f *File) Parse(in []byte) error {
	if len(in) != correctFileSize {
		return fmt.Errorf("expected size %d see %d", correctFileSize, len(in))
	}
	nameNullTerm := bytes.IndexByte(in[fileNameOffset:fileNameOffset+maxFileName], 0)
	f.Name = string(in[fileNameOffset : fileNameOffset+nameNullTerm])
	f.Orig = in
	f.Direction = in[partyDirectionOffset]
	f.X = in[partyXPos]
	f.Y = in[partyYPos]
	f.Food = binary.LittleEndian.Uint16(in[partyFoodOffset : partyGemOffset+2])
	f.BankGold = binary.LittleEndian.Uint32(in[bankGoldOffset : bankGoldOffset+4])
	f.BankGems = binary.LittleEndian.Uint32(in[bankGemOffset : bankGemOffset+4])
	f.Gold = binary.LittleEndian.Uint32(in[partyGoldOffset : partyGoldOffset+4])
	f.Gems = binary.LittleEndian.Uint32(in[partyGemOffset : partyGemOffset+4])
	var err error
	f.Characters, err = ParseCharacters(in)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) String() string {
	rv := ""
	rv += fmt.Sprintf("At x: %d y: %d facing: %s\n", f.X, f.Y, Directions[f.Direction])
	rv += fmt.Sprintf("Food: %d\n", f.Food)
	rv += fmt.Sprintf("Gold: %d (in bank: %d)\n", f.Gold, f.BankGold)
	rv += fmt.Sprintf("Gems: %d (in bank: %d)\n", f.Gems, f.BankGems)
	rv += fmt.Sprintf("Character (%d):\n", len(f.Characters))
	for _, char := range f.Characters {
		rv += fmt.Sprintf("%s\n", char)
	}
	return rv
}

func (f *File) WriteTo(w io.Writer) (n int64, err error) {
	written := int64(0)

	// write data before name
	wrote, err := w.Write(f.Orig[0:fileNameOffset])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write name
	nameBytes := make([]byte, maxFileName)
	copy(nameBytes, []byte("Default Characters"))
	copy(nameBytes, []byte(f.Name))
	// null teriminate it
	nameBytes[len(f.Name)] = 0
	wrote, err = w.Write(nameBytes)
	if err != nil {
		return written, err
	}

	// write data after name before first character
	wrote, err = w.Write(f.Orig[fileNameOffset+maxFileName : firstCharPos])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write characters
	for i := 0; i < numCharacters; i++ {
		wrote, err := f.Characters[i].WriteTo(w)
		written += int64(wrote)
		if err != nil {
			return written, err
		}
	}

	// write data after laster character before location
	lastCharPos := firstCharPos + (numCharacters * charSize)
	wrote, err = w.Write(f.Orig[lastCharPos:partyDirectionOffset])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write direction, x and y position
	wrote, err = w.Write([]byte{f.Direction, f.X, f.Y})
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write data after location before food
	wrote, err = w.Write(f.Orig[partyDirectionOffset+3 : partyFoodOffset])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write food
	foodBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(foodBytes, f.Food)
	wrote, err = w.Write(foodBytes)
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write data after food before bank gold
	wrote, err = w.Write(f.Orig[partyFoodOffset+2 : bankGoldOffset])
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write bank gold
	bankGoldBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bankGoldBytes, f.BankGold)
	wrote, err = w.Write(bankGoldBytes)
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write bank gems
	bankGemsBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bankGemsBytes, f.BankGems)
	wrote, err = w.Write(bankGemsBytes)
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write gold
	goldBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(goldBytes, f.Gold)
	wrote, err = w.Write(goldBytes)
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write gems
	gemsBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(gemsBytes, f.Gems)
	wrote, err = w.Write(gemsBytes)
	written += int64(wrote)
	if err != nil {
		return written, err
	}

	// write data after gems
	wrote, err = w.Write(f.Orig[partyGemOffset+4:])
	written += int64(wrote)
	return written, nil
}

func NewFile(path string) (*File, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	rv := &File{}
	err = rv.Parse(data)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

var Directions = map[byte]string{
	0: "North",
	1: "East",
	2: "South",
	3: "West",
}
