package main

import "testing"
import "strings"
import "fmt"

// Input mimik a csv Exel file.
var Input = `num point;x;y;z
5;793964,82;87556,226;210,6

`

func TestFixStringLen(t *testing.T) {
	const in, out, len = "4", "0004", 4
	if x := FixStringLen(in, len); x != out {
		t.Errorf("FixStringLen((%v, %v) = %v, want %v", in, len, x, out)
	}
}

func TestFixStringLen2(t *testing.T) {
	const in, out, len = "4", "0000000000000004", 16
	if x := FixStringLen(in, len); x != out {
		t.Errorf("FixStringLen((%v, %v) = %v, want %v", in, len, x, out)
	}
}

func TestFixLen(t *testing.T) {
	const in, out, len = 12, "0000000000000012", 16
	if x := FixLen(in, len); x != out {
		t.Errorf("FixLen((%v, %v) = %v, want %v", in, len, x, out)
	}
}

func TestRemoveComa(t *testing.T) {
	var ReadInput = strings.NewReader(Input)
	r := ExelCsvNewReader(ReadInput)
	rout, err := r.ReadAll()
	if err != nil { // check for error
		fmt.Println(err)
	}
	for ligne, tt := range rout {
		if ligne != 0 { // ignore first line
			if x := RemoveComa(tt[1:2]); x != "793964820" {
				t.Errorf("RemoveComa((%v) = %v, want %v", tt[1:2], x, "793964820")
			}
			if x := RemoveComa(tt[2:3]); x != "87556226" {
				t.Errorf("RemoveComa((%v) = %v, want %v", tt[2:3], x, "87556226")
			}
			if x := RemoveComa(tt[3:4]); x != "210600" {
				t.Errorf("RemoveComa((%v) = %v, want %v", tt[3:4], x, "210600")
			}
		}
	}
}
