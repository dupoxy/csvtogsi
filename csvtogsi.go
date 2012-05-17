// add licence

// csvtogsi for leica . Convert an Exel .csv file to a leica .gsi file.
package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// for debugging
const debug debugging = false // or flip to true

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
