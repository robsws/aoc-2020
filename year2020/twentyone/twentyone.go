package twentyone

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// PartOne - not yet implemented
func PartOne(filename string) string {
	nonAllergicCount, _ := analyseRecipes(filename)
	return fmt.Sprint(nonAllergicCount)
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	_, csvIngredients := analyseRecipes(filename)
	return csvIngredients
}

func analyseRecipes(filename string) (int, string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	re := regexp.MustCompile("^([a-z ]+) \\(contains ([a-z,\\s]+)\\)$")
	ingCandidateAllergens := make(map[string]utils.Set)
	recipes := make([]recipe, 0)
	// Store all possible ingredients and all possible candidate allergens
	for line := range fileStream {
		submatches := re.FindStringSubmatch(line)
		ingredients := strings.Split(submatches[1], " ")
		ingredientsSet := utils.MakeSetFromSlice(utils.StringToInterfaceSlice(ingredients))
		allergens := strings.Split(submatches[2], ", ")
		allergensSet := utils.MakeSetFromSlice(utils.StringToInterfaceSlice(allergens))
		recipes = append(recipes, recipe{ingredientsSet, allergensSet})
		for _, i := range ingredients {
			_, ok := ingCandidateAllergens[i]
			if !ok {
				ingCandidateAllergens[i] = utils.MakeSet()
			}
			ingCandidateAllergens[i].Union(allergensSet)
		}
	}
	// For each recipe, rule out allergens for all ingredients not on the list
	for _, r := range recipes {
		for ing := range ingCandidateAllergens {
			if !r.ingredients.Contains(ing) {
				for _, allergen := range r.allergens.ToSlice() {
					ingCandidateAllergens[ing].Remove(allergen)
				}
			}
		}
	}
	// Loop over ingredients list and deduce ingredient to allergen mapping
	allergenIngredient := make(map[string]string)
	nonAllergicIngredients := utils.MakeSet()
	for len(allergenIngredient)+nonAllergicIngredients.Len() < len(ingCandidateAllergens) {
		for ingredient := range ingCandidateAllergens {
			if ingCandidateAllergens[ingredient].Len() == 0 {
				nonAllergicIngredients.Add(ingredient)
				continue
			}
			if ingCandidateAllergens[ingredient].Len() == 1 {
				allergen := ingCandidateAllergens[ingredient].ToSlice()[0].(string)
				allergenIngredient[allergen] = ingredient
				for ingredient2 := range ingCandidateAllergens {
					if ingredient != ingredient2 {
						ingCandidateAllergens[ingredient2].Remove(allergen)
					}
				}
			}
		}
	}
	// Figure out how often the non-allergic ingredients appear
	totalNonAllergic := 0
	for _, r := range recipes {
		for _, ing := range nonAllergicIngredients.ToSlice() {
			if r.ingredients.Contains(ing) {
				totalNonAllergic++
			}
		}
	}
	// Construct comma separated list of allergic ingredients by allergen in alphabetical order
	allergens := make([]string, 0, len(allergenIngredient))
	for allergen := range allergenIngredient {
		allergens = append(allergens, allergen)
	}
	sort.Strings(allergens)
	ingredientsSorted := make([]string, len(allergens))
	for i, allergen := range allergens {
		ingredientsSorted[i] = allergenIngredient[allergen]
	}
	allergicCSV := strings.Join(ingredientsSorted, ",")
	return totalNonAllergic, allergicCSV
}

type recipe struct {
	ingredients utils.Set
	allergens   utils.Set
}
