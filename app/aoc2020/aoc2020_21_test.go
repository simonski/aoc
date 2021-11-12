package aoc2020

import (
	"fmt"
	"os"
	"sort"
	"testing"
)

func Test_AOC2020_21_Test(t *testing.T) {
	DATA := DAY_21_DATA
	// DATA := DAY_21_TEST_DATA
	db := NewFoodDB(DATA)
	if len(db.Food) != 4 {
		t.Errorf("Expected 4 foods, got %v\n", len(db.Food))
	}

	food := db.Food["0"]
	if food.NumberOfIngredients() != 4 {
		t.Errorf("food 0 should contain 4 ingredients but contains %v.\n", food.NumberOfIngredients())
	}
	if !food.ContainsIngredient("mxmxvkd") {
		t.Errorf("food 0 should contain ingredient %v but does not.\n", "mxmxvkd")
	}
	if !food.ContainsIngredient("kfcds") {
		t.Errorf("food 0 should contain ingredient %v but does not.\n", "kfcds")
	}
	if !food.ContainsIngredient("sqjhc") {
		t.Errorf("food 0 should contain ingredient %v but does not.\n", "sqjhc")
	}
	if !food.ContainsIngredient("nhms") {
		t.Errorf("food 0 should contain ingredient %v but does not.\n", "nhms")
	}

	if food.NumberOfAllergens() != 2 {
		t.Errorf("food 0 should contain 2 allergens but contains %v.\n", food.NumberOfAllergens())
	}
	if !food.ContainsAllergen("dairy") {
		t.Errorf("food 0 should contain allergen %v but does not.\n", "dairy")
	}
	if !food.ContainsAllergen("fish") {
		t.Errorf("food 0 should contain allergen %v but does not.\n", "fish")
	}

	for _, ingredient := range db.Ingredients {
		t.Errorf("Ingredient '%v' is present in %v foods.\n", ingredient.Value, len(ingredient.Foods))

	}
	t.Errorf("%v\n", db.Debug())
	ingredientAllergyMap := db.BuildIngredientAllergyMap()
	fmt.Printf("\n\n")
	allergyMap := make(map[string]string)
	allergiesToSort := make([]string, 0)
	for key, value := range ingredientAllergyMap {
		allergyMap[value] = key
		fmt.Printf("%v %v\n", value, key)
		allergiesToSort = append(allergiesToSort, value)
	}

	sort.Strings(allergiesToSort)
	fmt.Printf("\n\n%v\n\n", allergiesToSort)
	fmt.Printf("\n\n")

	line := ""
	for _, allergy := range allergiesToSort {
		fmt.Printf("%v - %v\n", allergy, allergyMap[allergy])
		line += allergyMap[allergy] + ","
	}
	fmt.Printf("LINE IS '%v'\n", line)
	// fmt.Printf("%v", ingredientAllergyMap)

	ingredientsWithoutAllergy := make([]string, 0)
	for _, ingredient := range db.Ingredients {
		_, exists := ingredientAllergyMap[ingredient.Value]
		if !exists {
			ingredientsWithoutAllergy = append(ingredientsWithoutAllergy, ingredient.Value)
		}
	}

	counter := 0
	for _, food := range db.Food {
		for _, ingredient := range ingredientsWithoutAllergy {
			if food.ContainsIngredient(ingredient) {
				counter++
			}
		}
	}
	fmt.Printf("Ingredients wihtout allergy occurs %v times.\n", counter)

	// fmt.Printf("Ingredients without allergy: %v\n", ingredientsWithoutAllergy)

	// for key, ingredient := range db.Ingredients {
	// 	_, isAllergenIngredient = db.Ingredients[]

	// }

	os.Exit(10)
	db.Analyse()
	t.Errorf("%v\n", db.Debug())

	db2 := NewFoodDB(DATA)
	counter = 0
	for ingredient, _ := range db.NoAllergens {
		for _, food := range db2.Food {
			if food.ContainsIngredient(ingredient) {
				counter++
			}
		}
	}
	fmt.Printf("There are %v no-allergen ingredients,they appear a total of %v times.\n", len(db.NoAllergens), counter)

	// ingredients := db.GetIngredientsWithoutAllergens()
	// if len(ingredients) != 5 {
	// 	t.Errorf("Ingredients without allergens should be 5 but is %v.\n", len(ingredients))
	// }

	if len(db.IdentifiedAllergens) > 0 {
		for k, v := range db.IdentifiedAllergens {
			fmt.Printf(" allergens: %v = %v\n", k, v)
		}
	} else {
		fmt.Printf(" No identified ingredient/allergenm relations.\n")
	}

}
