package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DBReader interface {
	Read(file string) Recipes
	Convert() []byte
}

type Ingredients struct {
	IngredientName  string `json:"ingredient_name" xml:"itemname"`
	IngredientCount string `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
}
type Cake struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}

type Recipes struct {
	Cake []Cake `json:"cake" xml:"cake"`
}

type JSONReader Recipes
type XMLReader Recipes

func (j *JSONReader) Read(file string) Recipes {
	jsonFile, _ := os.ReadFile(file)
	err := json.Unmarshal([]byte(jsonFile), &j)
	if err != nil {
		panic(err)
	}
	return Recipes(*j)
}

func (j *JSONReader) Convert() []byte {

	data, err := xml.MarshalIndent(j, "", "    ")

	if err != nil {
		panic(err)
	}

	return data
}

func (x *XMLReader) Read(file string) Recipes {
	jsonFile, _ := os.ReadFile(file)
	err := xml.Unmarshal([]byte(jsonFile), &x)

	if err != nil {
		panic(err)
	}
	return Recipes(*x)
}

func (x *XMLReader) Convert() []byte {

	data, err := json.MarshalIndent(x, "", "    ")

	if err != nil {
		panic(err)
	}

	return data
}

func main() {
	Start()
}

func Start() {

	FlagF := flag.String("f", "", "")
	flag.Parse()

	if strings.HasSuffix(*FlagF, ".xml") {
		var x XMLReader
		Parse(*FlagF, &x)
	}
	if strings.HasSuffix(*FlagF, ".json") {
		var j JSONReader
		Parse(*FlagF, &j)
	}
}

func Parse(file string, reader DBReader) {
	path := filepath.Join("../", "files/", file)
	fmt.Printf("%v\n\n%v", reader.Read(path), string(reader.Convert()))
}
