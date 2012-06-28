// Copyright 2011 The csvtogsi Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// csvtogsi for leica . Convert an Exel .csv file to a leica .gsi file.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	file := flag.String("file", "test.csv", "Set the file to import")
	flag.Parse()
	// ReadInput .
	Input, err := ioutil.ReadFile(*file)
	if err != nil { // check for error
		fmt.Println(err)
	}
	r := ExelCsvNewReader(bytes.NewReader(Input))
	// func (r *Reader) ReadAll() (records [][]string, err error)
	out, err := r.ReadAll()
	if err != nil { // check for error
		fmt.Println(err)
	}
	// fmt.Println(out)
	for i, tt := range out {
		if i != 0 { // ignore first line
			ligne := i
			out := fmt.Sprint("*11", FixLen(ligne, 4), // '*110001'
				"+", FixStringLen(RemoveComa(tt[:1]), 16), // '+0000000000005001'
				" 81..10+", FixStringLen(RemoveComa(tt[1:2]), 16), // ' 81..10+0000000793905635'
				" 82..10+", FixStringLen(RemoveComa(tt[2:3]), 16), // ' 82..10+0000000087528582'
				" 83..10+", FixStringLen(RemoveComa(tt[3:4]), 16)) // ' 83..10+0000000000210660'
			fmt.Println(out)
		}
	}
}
