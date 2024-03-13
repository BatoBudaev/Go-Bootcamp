package recipes

import "fmt"

func findCakeByName(cakes []Cake, name string) (Cake, bool) {
	for _, cake := range cakes {
		if cake.Name == name {
			return cake, true
		}
	}
	return Cake{}, false
}

func CompareDB(oldDB, newDB Recipes) {

	for _, oldCake := range oldDB.Cake {
		newCake, found := findCakeByName(newDB.Cake, oldCake.Name)
		if !found {
			fmt.Printf("REMOVED cake \"%s\"\n", oldCake.Name)
			continue
		}

		if oldCake.Time != newCake.Time {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldCake.Name, newCake.Time, oldCake.Time)
		}

		compareIngredients(oldCake.Ingredients, newCake.Ingredients, oldCake.Name)
	}

	for _, newCake := range newDB.Cake {
		_, found := findCakeByName(oldDB.Cake, newCake.Name)
		if !found {
			fmt.Printf("ADDED cake \"%s\"\n", newCake.Name)
		}
	}
}

func compareIngredients(oldIngredients, newIngredients []Ingredient, cakeName string) {
	oldIngredientMap := make(map[string]Ingredient)
	for _, oldIngredient := range oldIngredients {
		oldIngredientMap[oldIngredient.Name] = oldIngredient
	}

	newIngredientMap := make(map[string]Ingredient)
	for _, newIngredient := range newIngredients {
		newIngredientMap[newIngredient.Name] = newIngredient
	}

	for name, oldIngredient := range oldIngredientMap {
		newIngredient, found := newIngredientMap[name]
		if !found {
			fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\"\n", name, cakeName)
			continue
		}

		if oldIngredient.Count != newIngredient.Count {
			fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%v\" instead of \"%v\"\n",
				name, cakeName, newIngredient.Count, oldIngredient.Count)
		}

		if oldIngredient.Unit != newIngredient.Unit {
			fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
				name, cakeName, newIngredient.Unit, oldIngredient.Unit)
		}
	}

	for name := range newIngredientMap {
		_, found := oldIngredientMap[name]
		if !found {
			fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n", name, cakeName)
		}
	}
}
