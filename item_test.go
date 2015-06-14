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
	"reflect"
	"testing"
)

func TestItemString(t *testing.T) {
	tests := []struct {
		item     *Item
		expected string
	}{
		{
			item:     &Item{},
			expected: "",
		},
		{
			item: &Item{
				Elemental: 2,
				Substance: 22,
				Attribute: 9,
				Type:      5,
				Special:   4,
			},
			expected: "Fiery Obsidian Dragon Cutlass of Arrows",
		},
	}

	for _, test := range tests {
		actual := test.item.String()
		if actual != test.expected {
			t.Errorf("expected '%s' got '%s' for %#v", test.expected, actual, test.item)
		}
	}
}

func TestParseItem(t *testing.T) {
	tests := []struct {
		in   string
		item *Item
		err  error
	}{
		{
			in: "fiery obsidian dragon cutlass of arrows",
			item: &Item{
				Elemental: 2,
				Substance: 22,
				Attribute: 9,
				Type:      5,
				Special:   4,
			},
			err: nil,
		},
		{
			in: "obsidian dragon cutlass of arrows",
			item: &Item{
				Elemental: 0,
				Substance: 22,
				Attribute: 9,
				Type:      5,
				Special:   4,
			},
			err: nil,
		},
		{
			in: "dragon cutlass of arrows",
			item: &Item{
				Elemental: 0,
				Substance: 0,
				Attribute: 9,
				Type:      5,
				Special:   4,
			},
			err: nil,
		},
		{
			in: "cutlass of arrows",
			item: &Item{
				Elemental: 0,
				Substance: 0,
				Attribute: 0,
				Type:      5,
				Special:   4,
			},
			err: nil,
		},
		{
			in: "cutlass",
			item: &Item{
				Elemental: 0,
				Substance: 0,
				Attribute: 0,
				Type:      5,
				Special:   0,
			},
			err: nil,
		},
		{
			in: "cat",
			item: &Item{
				Elemental: 0,
				Substance: 0,
				Attribute: 0,
				Type:      0,
				Special:   0,
			},
			err: UnknownItemErr,
		},
		{
			in: "fiery obsidian dragon great_axe of arrows",
			item: &Item{
				Elemental: 2,
				Substance: 22,
				Attribute: 9,
				Type:      29,
				Special:   4,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		actual, err := NewItem(test.in)
		if err != test.err {
			t.Errorf("expected error %v, got %s", test.err, err)
			continue
		}
		if actual != nil && !reflect.DeepEqual(actual, test.item) {
			t.Errorf("expected %s\ngot %s", test.item.String(), actual.String())
		}
	}
}
