// Copyright 2011 The csvtogsi Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
)

// for debugging
const debug debugging = true // or flip to false

type debugging bool

func (d debugging) Printf(format string, args ...interface{}) {
	if d {
		log.Printf(format, args...)
	}
}

// ExelCsvNewReader returns a new csv.Reader that reads from r with comma cet to ';'.
func ExelCsvNewReader(r io.Reader) *csv.Reader {
	// func NewReader(r io.Reader) *Reader
	set := csv.NewReader(r)
	set.Comma = ';' // set Comma to ;
	return set
}

// RemoveComa returns a string that reads from slice with added 0 digit to always have 'X,XXX'.
func RemoveComa(slice []string) string {
	if strings.Contains(slice[0], ",") == true {
		c := strings.SplitAfter(slice[0], ",")
		for len(c[1]) < 3 {
			// add 0 to end of string
			s1 := c[1] + strings.Repeat("0", 3-len(c[1]))
			out := strings.Replace(c[0], ",", "", 1) + s1
			debug.Printf("RemoveComa if for: slice[0] %v ; c[0] %v ; c[1] %v ; s1 %v ; out %v ", slice[0], c[0], c[1], s1, out)
			return out
		}
	}
	out := strings.Replace(slice[0], ",", "", 1)
	debug.Printf("RemoveComa: slice[0] %v ; out %v ", slice[0], out)
	return out
}

// FixLen returns a string that reads from i with added ... digit to the start to always have x digit.
func FixLen(i int, x int) string {
	s := strconv.Itoa(i)
	out := FixStringLen(s, x)
	debug.Printf("FixLen: %v ", out)
	return out

}

// FixStringLen returns a string that reads from slice with added ... digit to the start to always have x digit.
func FixStringLen(slice string, x int) string {
	if len(slice) < x {
		for len(slice) < x {
			// add 0 to start of string
			s1 := strings.Repeat("0", x-len(slice)) + slice
			out := s1
			debug.Printf("FixStringLen if for: %v ", out)
			return out

		}
	}
	out := slice
	debug.Printf("FixStringLen: %v ", out)
	return out

}
