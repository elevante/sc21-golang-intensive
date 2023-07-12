package main

type Ingredients struct {
	Ingredient_name  string `json:"ingredient_name" xml:"itemname"`
	Ingredient_count string `json:"ingredient_count" xml:"itemcount"`
	Ingredient_unit  string `json:"ingredient_unit" xml:"itemunit,omitempty"`
}
type Cakes struct {
	Name       string `json:"name" xml:"name"`
	Time       string `json:"time" xml:"stovetime"`
	Ingredient []Ingredients
}

type Recipes struct {
	Cake []Cakes `json:"cake" xml:"cake"`
}

func main() {

}
