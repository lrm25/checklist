package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

const DATA_LOCATION = "procedures.json"
const EXIT_STRING = "Exit"

type ProcedureElement struct {
	Name     string
	Location string
}

type ProcedureElements struct {
	Procedures []ProcedureElement
}

type Data struct {
	data         ProcedureElements
	indices      []int
	procedureMap map[int]ProcedureElement
}

func NewData() *Data {
	return &Data{
		data:         ProcedureElements{},
		indices:      []int{},
		procedureMap: make(map[int]ProcedureElement),
	}
}

func (d *Data) Load() error {
	jsonFile, err := os.Open(DATA_LOCATION)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(jsonBytes, &d.data); err != nil {
		return err
	}

	index := 1
	for _, p := range d.data.Procedures {
		d.procedureMap[index] = p
		d.indices = append(d.indices, index)
		index++
	}
	d.procedureMap[index] = ProcedureElement{Name: "Exit", Location: ""}
	d.indices = append(d.indices, index)

	return nil
}

func (d *Data) Select() (string, error) {

	sort.Ints(d.indices)

	println("Select procedure:")
	for _, i := range d.indices {
		println(fmt.Sprintf("%d: "+d.procedureMap[i].Name, i))
	}

	validSelection := false
	indicesCount := len(d.indices)
	var selection int
	for !validSelection {
		print("Select value: ")
		if _, err := fmt.Scan(&selection); err != nil {
			println(err.Error())
			continue
		}
		if selection < 1 || indicesCount < selection {
			fmt.Printf("Select value with minimum of 1 and maximum of %d\n", indicesCount)
			continue
		}
		if d.procedureMap[selection].Name == "Exit" {
			os.Exit(0)
		}
		validSelection = true
	}
	return d.procedureMap[selection].Location, nil
}
