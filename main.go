package main

import (
	"github.com/lrm25/checklist/data"
	"github.com/lrm25/checklist/filecopy"
)

func main() {
	d := data.NewData()
	if err := d.Load(); err != nil {
		println("Error loading JSON data: " + err.Error())
		return
	}
	location, err := d.Select()
	if err != nil {
		println("Error selecting option: " + err.Error())
		return
	}

	copy := filecopy.NewFileCopy(location)
	if err := copy.Create(); err != nil {
		println("Error creating file copy:  " + err.Error())
		return
	}
	if err := copy.Open(); err != nil {
		println("Error opening file: " + err.Error())
	}
}
