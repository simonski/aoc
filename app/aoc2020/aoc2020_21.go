package aoc2020

/*
--- Day 21: Allergen Assessment ---
You reach the train's last stop and the closest you can get to your vacation island without getting wet. There aren't even any boats here, but nothing can stop you now: you build a raft. You just need a few days' worth of food for your journey.

You don't speak the local language, so you can't read any ingredients lists. However, sometimes, allergens are listed in a language you do understand. You should be able to use this information to determine which ingredient contains which allergen and work out which foods are safe to take with you on your trip.

You start by compiling a list of foods (your puzzle input), one food per line. Each line includes that food's ingredients list followed by some or all of the allergens the food contains.

Each allergen is found in exactly one ingredient. Each ingredient contains zero or one allergen. Allergens aren't always marked; when they're listed (as in (contains nuts, shellfish) after an ingredients list), the ingredient that contains each listed allergen will be somewhere in the corresponding ingredients list. However, even if an allergen isn't listed, the ingredient that contains that allergen could still be present: maybe they forgot to label it, or maybe it was labeled in a language you don't know.

For example, consider the following list of foods:

mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
The first food in the list has four ingredients (written in a language you don't understand): mxmxvkd, kfcds, sqjhc, and nhms. While the food might contain other allergens, a few allergens the food definitely contains are listed afterward: dairy and fish.

The first step is to determine which ingredients can't possibly contain any of the allergens in any food in your list. In the above example, none of the ingredients kfcds, nhms, sbzzf, or trh can contain an allergen. Counting the number of times any of these ingredients appear in any ingredients list produces 5: they all appear once each except sbzzf, which appears twice.

Determine which ingredients cannot possibly contain any of the allergens in your list. How many times do any of those ingredients appear?
*/

import (
	"fmt"
	"strings"

	cli "github.com/simonski/cli"
)

// AOC_2020_20 is the entrypoint
func AOC_2020_21(cli *cli.CLI) {
	AOC_2020_21_part1_attempt1(cli)
}

func AOC_2020_21_part1_attempt1(cli *cli.CLI) {
}

func NewFoodDB(data string) *FoodDB {
	splits := strings.Split(data, "\n")
	db := FoodDB{}
	db.Food = make(map[string]*Food)
	db.Ingredients = make(map[string]*Ingredient)
	db.Allergens = make(map[string]*Allergen)

	for index, line := range splits {
		db.AddFood(index, line)
	}
	return &db
}

type FoodDB struct {
	Food        map[string]*Food
	Ingredients map[string]*Ingredient
	Allergens   map[string]*Allergen

	NoAllergens         map[string]bool
	IdentifiedAllergens map[string]string
}

func (db *FoodDB) Debug() string {
	line := fmt.Sprintf("DB: %v food, %v ingredients, %v allergens.\n", len(db.Food), len(db.Ingredients), len(db.Allergens))
	line += fmt.Sprintf("No allergens: %v, IdentifiedAllergens: %v\n", db.NoAllergens, db.IdentifiedAllergens)
	return line
}

func (db *FoodDB) GetIngredientsWithoutAllergens() []*Ingredient {
	ingredients := make([]*Ingredient, 0)
	return ingredients
}

func (db *FoodDB) BuildIngredientAllergyMap() map[string]string {

	shared := make(map[string][]string)
	for _, allergen := range db.Allergens {
		sharedIngredients := allergen.CreateSharedIngredientsAcrossFoods()
		// remove these ingredients from everything else
		shared[allergen.Value] = sharedIngredients

	}

	// for key, value := range shared {
	// 	fmt.Printf("%v = %v\n", key, value)
	// }

	for {
		quit := true
		for allergen, ingredients := range shared {
			if len(ingredients) == 1 {
				ingredient := ingredients[0]
				for allergen2, ingredients2 := range shared {
					if allergen2 == allergen {
						continue
					}
					arr := make([]string, 0)
					for _, i := range ingredients2 {
						if i != ingredient {
							arr = append(arr, i)
						} else {
							quit = false
						}
						shared[allergen2] = arr
					}
				}
			}
		}
		// fmt.Printf("\n")
		// for key, value := range shared {
		// 	fmt.Printf("%v = %v\n", key, value)
		// }
		if quit {
			break
		}
	}

	s := make(map[string]string)
	for key, value := range shared {
		s[value[0]] = key
	}
	return s

}

func (db *FoodDB) Analyse() {
	/*

		   trh fvjkl sbzzf mxmxvkd (contains dairy)
		   mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
		   sqjhc fvjkl (contains soy)
		   sqjhc mxmxvkd sbzzf (contains fish)

		   for each ingredient
			   if total_occurrances is 1
				   put in does not occur anywhere
				   remove from food
				   if food.ingredients == 1 and food.allergens == 1:
					   # this allergen and ingredient are identified
					   create allergen/ingredient
					   # remove them both, remove the food



		  Step 1 Remove all 1-occurances that we can
		   kfcds nhms sbzzf, trh

		   trh fvjkl sbzzf mxmxvkd (contains dairy)
		   mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
		   sqjhc fvjkl (contains soy)
		   sqjhc mxmxvkd sbzzf (contains fish)

		Step 2

		   ---- fvjkl ----- mxmxvkd (contains dairy)
		   mxmxvkd ---- sqjhc ---- (contains dairy, fish)
		   sqjhc fvjkl (contains soy)
		   sqjhc mxmxvkd ----- (contains fish)

		// all ingrediens are in > 1 food now
		   // can we attach an ingredient to an allergen
		   // take ingredient 1 'fvjkl'

		   for an allergen 'dairy'
		    list foods, strip other allergens for this purpose
			   ---- fvjkl ----- mxmxvkd (contains dairy)
			   mxmxvkd ---- sqjhc ---- (contains dairy, ----)

			remove ingredients NOT in both

			   ---- ------ ----- mxmxvkd (contains dairy)
			   mxmxvkd ---- ------ ---- (contains dairy, ---)

			mxmxvkd is dairy
			remove allergen dairy from everywhere
			remove ingredient mxmxvkd from everywhere
			IDENTIFY identify mxmxvkd as owner of allergen dairy






	*/

	noAllergens := make(map[string]bool)
	identifiedAllergens := make(map[string]string)
	iteration := 0
	for {
		quit := true
		// Step 1
		fmt.Printf("\n[Iteration %v Step 1]\n", iteration)
		noAllergensThisIteration := 0
		singleFoodIngredientsThisIteration := 0
		for _, ingredient := range db.Ingredients {
			if len(ingredient.Foods) == 1 {
				// it is present in only 1 food
				singleFoodIngredientsThisIteration++
				for _, food := range ingredient.Foods {
					// check the allergen list, if the allergens exist elsewhere, this ingredient has no allergens
					for _, allergen := range food.Allergens {
						if len(allergen.Foods) > 1 {
							// then this ingredient has no allergens
							fmt.Printf("[Iteration %v Step 1: Ingredient '%v' can be removed as a 'no-allergen' ingredient.\n", iteration, ingredient.Value)
							noAllergens[ingredient.Value] = true
							db.RemoveIngredient(ingredient.Value)
							noAllergensThisIteration++
							// we made a change, so we need to walk our database once more
							quit = false
						}
					}
				}
			}
		}

		if singleFoodIngredientsThisIteration == 0 {
			fmt.Printf("[Iteration %v Step 1] No ingredients are in a single food, all are in multiple.\n", iteration)
		} else {
			fmt.Printf("[Iteration %v Step 1] %v ingredients were in a single food, %v were no-allergens and removed..\n", iteration, singleFoodIngredientsThisIteration, noAllergensThisIteration)
		}
		/*
			// Step 2
			---- fvjkl ----- mxmxvkd (contains dairy)
			mxmxvkd ---- sqjhc ---- (contains dairy, fish)
			sqjhc fvjkl (contains soy)
			sqjhc mxmxvkd ----- (contains fish)

			all ingredients are in > 1 food now
			now go over allergens themselves
			first allergen 'dairy', list foods
				---- fvjkl ----- mxmxvkd (contains dairy)
				mxmxvkd ---- sqjhc ---- (contains dairy, fish)

				for dairy, the common ingredient is mxmxvkd
				so mxmxvkd is dairy
				add this to the ActualAllergens, then remove mxmxvkd ingredient AND 'dairy' allergen
						quit = false

				for each food
					if number of allergens == 0
						each ingredeint in this food is a zero allergen ingredient
						ADD to NO ALLERGENS
						REMOVE INGREDIENT
						quit = false
		*/

		fmt.Printf("[Iteration %v Step 2]\n", iteration)
		for _, allergen := range db.Allergens {
			fmt.Printf("[Iteration %v Step 2] allergen %v has %v foods.\n", iteration, allergen.Value, len(allergen.Foods))
			// if len(allergen.Foods) > 1 {
			sharedIngredients := allergen.CreateSharedIngredientsAcrossFoods()
			fmt.Printf("[Iteration %v Step 2] Shared Ingredients for allergen %v : %v\n", iteration, allergen.Value, sharedIngredients)
			if len(sharedIngredients) == 1 {
				// this is the only common ingredient, this must be the ingredient to use
				ingredient := sharedIngredients[0]
				identifiedAllergens[ingredient] = allergen.Value
				db.RemoveIngredient(ingredient)
				db.RemoveAllergen(allergen.Value)
				quit = false
			}
			// this ingredient appears in > 1 allergen
			// remove any ingredients that are NOT in both for this allergen as they cannot be it
			// }
		}

		for _, food := range db.Food {
			if len(food.Allergens) == 0 {
				// each ingredeint in this food is a zero allergen ingredient
				for _, ingredient := range food.Ingredients {
					noAllergens[ingredient.Value] = true
					db.RemoveIngredient(ingredient.Value)
					quit = false
					// ADD to NO ALLERGENS
					// REMOVE INGREDIENT
					// quit = false
				}
			}
		}

		// for an allergen 'dairy'
		//  list foods, strip other allergens for this purpose
		// 	---- fvjkl ----- mxmxvkd (contains dairy)
		// 	mxmxvkd ---- sqjhc ---- (contains dairy, ----)

		//  remove ingredients NOT in both

		// 	---- ------ ----- mxmxvkd (contains dairy)
		// 	mxmxvkd ---- ------ ---- (contains dairy, ---)

		//  mxmxvkd is dairy
		//  remove allergen dairy from everywhere
		//  remove ingredient mxmxvkd from everywhere
		//  IDENTIFY identify mxmxvkd as owner of allergen dairy
		if quit {
			fmt.Printf("Quitting.\n")
			break
		}
		iteration++
	}

	// Now we have removed the ingredients

	db.NoAllergens = noAllergens
	db.IdentifiedAllergens = identifiedAllergens
}

func (db *FoodDB) AddFood(index int, line string) *Food {
	line = strings.ReplaceAll(line, "(", "")
	line = strings.ReplaceAll(line, ")", "")
	line = strings.ReplaceAll(line, ",", "")
	splits := strings.Split(line, " ")
	isIngredients := true
	key := fmt.Sprintf("%v", index)
	food := &Food{Key: key, Value: line}
	food.Ingredients = make(map[string]*Ingredient)
	food.Allergens = make(map[string]*Allergen)
	db.Food[key] = food

	for _, value := range splits {
		if value == "contains" {
			isIngredients = false
			continue
		}
		if isIngredients {
			ingredient := db.GetIngredient(value)
			food.AddIngredient(ingredient)
			ingredient.AddFood(food)
		} else {
			allergen := db.GetAllergen(value)
			food.AddAllergen(allergen)
			allergen.AddFood(food)
		}
	}
	return food
}

func (db *FoodDB) RemoveIngredient(value string) {
	delete(db.Ingredients, value)
	for _, a := range db.Allergens {
		a.RemoveIngredient(value)
	}
	for _, f := range db.Food {
		f.RemoveIngredient(value)
	}
}

func (db *FoodDB) RemoveAllergen(value string) {
	delete(db.Allergens, value)
	for _, i := range db.Ingredients {
		i.RemoveAllergen(value)
	}
	for _, f := range db.Food {
		f.RemoveAllergen(value)
	}
}

func (db *FoodDB) GetIngredient(value string) *Ingredient {
	ingredient, exists := db.Ingredients[value]
	if exists {
		return ingredient
	}
	ingredient = &Ingredient{Value: value}
	ingredient.Allergens = make(map[string]*Allergen)
	ingredient.Foods = make(map[string]*Food)
	db.Ingredients[value] = ingredient
	return ingredient
}

func (db *FoodDB) GetAllergen(value string) *Allergen {
	allergen, exists := db.Allergens[value]
	if exists {
		return allergen
	}
	allergen = &Allergen{Value: value}
	allergen.Ingredients = make(map[string]*Ingredient)
	allergen.Foods = make(map[string]*Food)
	db.Allergens[value] = allergen
	return allergen
}

type Food struct {
	Key   string
	Value string

	Ingredients map[string]*Ingredient
	Allergens   map[string]*Allergen
}

func (food *Food) ContainsIngredient(value string) bool {
	for _, v := range food.Ingredients {
		if v.Value == value {
			return true
		}
	}
	return false
	// _, exists := food.Ingredients[value]
	// return exists
}

func (food *Food) ContainsAllergen(value string) bool {
	for _, v := range food.Allergens {
		if v.Value == value {
			return true
		}
	}
	return false
	// _, exists := food.Allergens[value]
	// return exists
}

func (food *Food) NumberOfIngredients() int {
	return len(food.Ingredients)
}

func (food *Food) NumberOfAllergens() int {
	return len(food.Allergens)
}

func (food *Food) AddIngredient(ingredient *Ingredient) {
	food.Ingredients[ingredient.Value] = ingredient
}

func (food *Food) RemoveIngredient(value string) {
	delete(food.Ingredients, value)
}

func (food *Food) AddAllergen(allergen *Allergen) {
	food.Allergens[allergen.Value] = allergen
}

func (food *Food) RemoveAllergen(value string) {
	delete(food.Allergens, value)
}

type Ingredient struct {
	Value     string
	Foods     map[string]*Food
	Allergens map[string]*Allergen
}

func (ingredient *Ingredient) AddFood(food *Food) {
	ingredient.Foods[food.Value] = food
}

func (ingredient *Ingredient) AddAllergen(allergen *Allergen) {
	ingredient.Allergens[allergen.Value] = allergen
}

func (ingredient *Ingredient) RemoveAllergen(value string) {
	delete(ingredient.Allergens, value)
}

type Allergen struct {
	Value       string
	Foods       map[string]*Food
	Ingredients map[string]*Ingredient
}

func (allergen *Allergen) AddFood(food *Food) {
	allergen.Foods[food.Value] = food
}

func (allergen *Allergen) AddIngredient(ingredient *Ingredient) {
	allergen.Ingredients[ingredient.Value] = ingredient
}

func (allergen *Allergen) RemoveIngredient(value string) {
	delete(allergen.Ingredients, value)
}

func (allergen *Allergen) CreateSharedIngredientsAcrossFoods() []string {
	numOfFoods := len(allergen.Foods)
	sharedIngredients := make([]string, 0)
	countOfIngredients := make(map[string]int)
	for _, food := range allergen.Foods {
		for _, ingredient := range food.Ingredients {
			// fmt.Printf("Allergen '%v' has food '%v' which has ingredient '%v'\n", allergen.Value, food.Value, ingredient.Value)
			value, exists := countOfIngredients[ingredient.Value]
			if exists {
				value++
				countOfIngredients[ingredient.Value] = value
				// fmt.Printf("ingredient '%v' count is now '%v'\n", ingredient.Value, value)
			} else {
				value = 1
				countOfIngredients[ingredient.Value] = value
				// fmt.Printf("ingredeient '%v' count is now %v\n", ingredient.Value, value)
			}
		}
	}

	for k, v := range countOfIngredients {
		if v == numOfFoods {
			sharedIngredients = append(sharedIngredients, k)
		}
	}
	return sharedIngredients

}
