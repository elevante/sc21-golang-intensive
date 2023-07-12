package main

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

var jsonReader Recipes
var xmlReader Recipes

func main() {

}
