// Copyright 2011 The csvtogsi Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Leica Geo Serial Interface
//
// From Version 2.20 onwards, there is a choice of two
// GSI formats, with word lengths of 8 and 16 characters
// respectively. When 16 characters are stored and supported
// the following special conditions apply:
//	- A measurement block is tagged with * at the first position.
//	- A data word includes the data at position 7 to 23 instead of
//		at 7 to 15
// Word format :
//	|------------------------|
//	|84..10+0000123456789123 |
//	|WW....+0000123456789123 |
//	|--------position--------|
//	|123456789111111111122222|
//	|         012345678901234|
//	|------------------------|
//	|position---|meaning-----|
//	|-----------|------------|
//	|1-2        |WordIndex   | range 01-99
//	|3-6        |Additio.Info|
//	|7-15(23)   |Data        |
//	|16(24)     |Blank       |
//	|------------------------|
type Word struct {
	WordIndex      WordIndex          // 84
	AdditionalInfo AdditionalInfo     // ..10
	Data           Data               // +0000123456789123
	Blank          SeparatingCaracter // separating character

}

// Word Identification table :
//	|------------------------------------------|
//	|WordIndex--|Description-------------------|
//	|-----------|------------------------------|
//	|General    |                              |
//	|11         |Point number (includes        |
//	|           |block number)                 |
//	|12         |Instrument serial number      |
//	|13         |Instrument type               |
//	|18         |Time format 1:                |
//	|           |pos. 8-9   year,              |
//	|           |     10-11 sec,               |
//	|           |     12-14 msec               |
//	|19         |Time format 2:                |
//	|           |pos. 8-9   month,             |
//	|           |     10-11 day,               |
//	|           |     12-13 hour               |
//	|           |     14-15 min                |
//	|-----------|------------------------------|
//	|Angles     |                              |
//	|21         |Horizontal circle (Hz)        |
//	|22         |Vertical angle (V)            |
//	|25         |Horizontal circle difference  |
//	|           |(Hz0-Hz)                      |
//	|-----------|------------------------------|
//	|Distances  |                              |
//	|31         |Slope distance                |
//	|32         |Horizontal distance           |
//	|33         |Height difference             |
//	|-----------|------------------------------|
//	|Codebl.    |                              |
//	|41         |Code number (includes         |
//	|           |block number)                 |
//	|42-49      |Information 1-8               |
//	|-----------|------------------------------|
//	|Distance   |                              |
//	|additional |                              |
//	|information|                              |
//	|51         |Constants (ppm,mm)            |
//	|52         |Number of measurements        |
//	|53         |Standard deviation            |
//	|58         |Signal strength               |
//	|59         |Reflector constant            |
//	|           |(1/10 mm) ppm                 |
//	|-----------|------------------------------|
//	|Point      |                              |
//	|coding     |                              |
//	|71         |Point code                    |
//	|72-79      |Attrib. 1-8                   |
//	|-----------|------------------------------|
//	|Coordinates|                              |
//	|81         |Easting (target)              |
//	|82         |Northing (target)             |
//	|83         |Elevation (target)            |
//	|84         |Station easting (Eo)          |
//	|85         |Station northing (No)         |
//	|86         |Station elevation (Ho)        |
//	|87         |Reflector height(above ground)|
//	|88         |Instrumentheight(above ground)|
//	|-----------|------------------------------|
type WordIndex struct {
	WordIndex int // 84

}

// Additional Info
// Positions 3 to 6 contain information witch relates to the
// data in positions 7 to 15(23)
//	|----------------------------------------------|-------------|
//	|Position---|Explanation-----------------------|Applicable-to|
//	|in-word----|----------------------------------|-------------|
//	|-----------|----------------------------------|-------------|
//	|3          |Extension for word index          |Digital level|
//	|-----------|----------------------------------|-------------|
//	|4          |Compensator information           |All words    |
//	|           |0 Automatic height index and      |containing   |
//	|           |monitoring of levelling up:OFF    |angle info.  |
//	|           |3 Automatic height index and      |             |
//	|           |monitoring of levelling up:ON     |             |
//	|-----------|----------------------------------|-------------|
//	|5          |Input mode                        |All words    |
//	|           |0 Value measured automatically    |containing   |
//	|           |1 Manual input from keyboard      |measured     |
//	|           |2 Angle: All Hz-corrections       |data         |
//	|           |for Line of sight error, tilting  |             |
//	|           |axis and tilt of standing axis: ON|             |
//	|           |3 Angle: All Hz-corrections OFF   |             |
//	|           |4 Result calculated from functions|             |
//	|-----------|----------------------------------|-------------|
//	|6          |Units                             |All words    |
//	|           |0 Metre (last digit = 1mm)        |containing   |
//	|           |1 US-feet (last digit = 1/1000ft) |measured     |
//	|           |2 400gon                          |data         |
//	|           |3 360° decimal                   |             |
//	|           |4 360° sexagesimal               |             |
//	|           |5 6400 mil                        |             |
//	|           |6 Metre (last digit = 1/10mm)     |             |
//	|           |7 US-feet (last digit = 1/10000ft)|             |
//	|           |8 Metre (last digit = 1/100mm)    |             |
//	|-----------|----------------------------------|-------------|
//
// A point in any of the positions 3 to 6 means that there
// is no information.
//
// For the data words Point number (11) and Code number (41), the
// block number occupies positions 3 to 6.
type AdditionalInfo struct {
	ExtensionWI int
	Compensator int
	InputMode   int
	Units       int
}

// Data (Positions 7-15/23)
//	|------------------------------------------------|-------------|
//	|Position---|Explanation-------------------------|Applicable-to|
//	|in-word----|------------------------------------|-------------|
//	|-----------|------------------------------------|-------------|
//	|7          |Extension of word indentification   |All words    |
//	|-----------|------------------------------------|-------------|
//	|8-15(23)   |The data comprises 8(16)            |All words    |
//	|           |numeric or alphanumeric             |containing   |
//	|           |characters                          |data.        |
//	|           |                                    |             |
//	|           |Note that certain words can         |Words 51-59  |
//	|           |contain two data blocks, i.e. pairs.|             |
//	|           |These are transferred automatically,|             |
//	|           |complete with signs, from the survey|             |
//	|           |instrument:                         |             |
//	|           |e.g. 0123 -035                      |             |
//	|           |     ppm   mm                       |             |
//	|-----------|------------------------------------|-------------|
type Data struct {
	Extension string // +/-
	ThisData  string // 0000123456789123

}

// Separating character (Positions 16/24)
//	|------------------------------------------------|-------------|
//	|Position---|Explanation-------------------------|Applicable-to|
//	|in-word----|------------------------------------|-------------|
//	|-----------|------------------------------------|-------------|
//	|16(24)     |Blank (separating character)        |All words    |
//	|-----------|------------------------------------|-------------|
// The last data word of a block must also contain the separating character
// and CRLF
type SeparatingCaracter struct {
	Sep  string // " "
	Last bool   // true/false

}

// Block number :
// Each data block is allocated a block number by the recording device.
// Block numbers start at 1 and increment automatically
// The Block number is contained in the firts word of the block.
// The first word of a measurement block is the point number (WI = 11).
// The first word of a code block is the code number (WI = 41).
// The structure of the first word of a data block is as follows:
//	|------------------------------------------------|
//	|Position---|Explanation-------------------------|
//	|in-word----|------------------------------------|
//	|-----------|------------------------------------|
//	|1-2        |Word index 11 or 41                 |
//	|3-6        |Block number                        |
//	|7          |Sign + or -                         |
//	|8-15(23)   |Point number, code number oer text  |
//	|16(24)     |Blank = separating character        |
//	|-----------|------------------------------------|
type BlockNumber struct {
	WordIndex WordIndex          // 11
	Blocknum  int                // 0001
	Sign      string             // +
	Pointnum  string             // 0000000000000001
	Blank     SeparatingCaracter // separating character

}

// Data transmitted by instruments over the GSI interface is composed of
// blocks. Each one of these data blocks is treated as a whole, and ends
// with a terminator (CR, or CR LF). There are two types of data block:
//	1 Measurement blocks
//	2 Code blocks
//
// Measurement blocks contain a point number and measurement information.
// They are designed primarily for triangulation, traverse, detail, and
// tacheometric surveys.
//
// Code blocks are designed primarily for recording identification codes,
// data-processing codes and information. However, they can also be used
// for recording measurement information such as instrument and target
// heights, and tie distances.
//
// Each data block has a block number. The block numbers start with 1 and
// are incremented by 1 each time a new data block is stored.
//
// A data block consists of words, each with 16 (24) characters.
// The maximum number of words in the TPS1100 is 12.
//
// All words can be used in measurement blocks
// exept words 41-49, with are reserved.
//
// A code block begins with 41, the word index for a
// code number.
type Block struct {
	// "*110001+0000000000005001 81..10+0000000793905635 82..10+0000000087528582 83..10+0000000000210660  "
	ThisBlock []Word
}
