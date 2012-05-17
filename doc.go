/*
Setup GOPATH

on windows:
	set GOPATH=C:\...\mygsi
on linux:
	export GOPATH=$HOME/mygsi
	
To Install

	go get -u dupoxy.ouvaton.org/p/go/csvtogsi/
	go install dupoxy.ouvaton.org/p/go/csvtogsi/
	
To run

goto your GOPATH bin dir and run

on windows:
	csvtogsi.exe -file="test.csv"

on linux:
	csvtogsi -file="test.csv"

Nota Bene

expected csv file to look like this:

	num point;x;y;z
	5001;793905,635;87528,582;210,66
	5002;793964,169;87547,069;211,01
	1;793971,847;87572,175;210,74
	2;794020,145;87571,815;211,16
	3;794020,026;87555,815;211,16
	4;793964,827;87556,226;210,66

*/
package main
