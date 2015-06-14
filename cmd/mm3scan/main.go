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
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
)

var offset = flag.Int("offset", 0, "offset to skip over")
var strSearch = flag.String("string", "", "string to search for")
var uintSearch = flag.Int("uint", 0, "unsigned integer to search for")

func main() {

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("filename required\n")
		return
	}

	data, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Print(err)
		return
	}

	if *offset > len(data) {
		fmt.Printf("offset %d greater than data length %d\n", *offset, len(data))
	}

	if *offset > 0 {
		data = data[*offset:]
	}

	if len(*strSearch) > 0 {
		fmt.Printf("Searcing for '%s': ", *strSearch)
		pos := bytes.Index(data, []byte(*strSearch))
		if pos < 0 {
			fmt.Printf("not found\n")
		} else {
			fmt.Printf("found at pos %d (%x)\n", pos+*offset, pos+*offset)
		}
	}

	if *uintSearch > 0 {
		uintBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(uintBytes, uint64(*uintSearch))
		if uintBytes[7] == 0 && uintBytes[6] == 0 {
			uintBytes = uintBytes[:6]
		}
		if uintBytes[5] == 0 && uintBytes[4] == 0 {
			uintBytes = uintBytes[:4]
		}
		if uintBytes[3] == 0 && uintBytes[2] == 0 {
			uintBytes = uintBytes[:2]
		}
		fmt.Printf("Searching for %d as '% x': ", *uintSearch, uintBytes)
		pos := bytes.Index(data, uintBytes)
		if pos < 0 {
			fmt.Printf("not found\n")
		} else {
			fmt.Printf("found at pos %d (%x)\n", pos+*offset, pos+*offset)
		}
	}
}
