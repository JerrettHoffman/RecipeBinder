package mock

import (
	"RecipeBinder/internal"
	"log"
	"sort"
	"strings"
	"time"

	"errors"
)

type MockRecipeDb struct {
	recipes   map[internal.ID]internal.RecipeData
	currentId internal.ID
}

func (m *MockRecipeDb) ReadRecipe(recipeId internal.ID) (internal.RecipeData, error) {
	if recipe, ok := m.recipes[recipeId]; ok {
		log.Printf("Read recipe %d: %v", recipeId, recipe)
		return recipe, nil
	} else {
		return internal.RecipeData{}, errors.New("Recipe not in db")
	}
}

func (m *MockRecipeDb) UpdateRecipe(recipe internal.RecipeData, recipeId internal.ID) error {
	m.recipes[recipeId] = recipe
	log.Printf("Updated recipe %d: %v", recipeId, recipe)
	return nil
}

func (m *MockRecipeDb) CreateRecipe(recipe internal.RecipeData, uploaderId internal.ID) (internal.ID, error) {
	if m.recipes == nil {
		m.recipes = make(map[internal.ID]internal.RecipeData)
	}

	m.currentId++
	recipeId := m.currentId
	m.recipes[recipeId] = recipe
	log.Printf("Added recipe: %v", recipe)
	return recipeId, nil
}

func (m *MockRecipeDb) DeleteRecipe(recipeId internal.ID) error {
	delete(m.recipes, recipeId)
	log.Printf("Deleted recipe %d", recipeId)
	return nil
}

func stringsMatch(query, other string) bool {
	return query == "" || strings.EqualFold(query, other)
}

func (m *MockRecipeDb) Search(params internal.SearchParams) []internal.SearchResult {
	log.Printf("Searching for %v", params)
	result := make([]internal.SearchResult, 0, 8)
	for id, recipe := range m.recipes {
		match := stringsMatch(params.RecipeName, recipe.RecipeName) &&
			stringsMatch(params.AuthorName, recipe.Author) &&
			stringsMatch(params.UploaderName, recipe.Uploader) &&
			// 1s is the zero value for duration
			(params.PrepTime == time.Second || params.PrepTime == recipe.PrepTime) &&
			(params.PrepTime == time.Second || params.TotalTime == recipe.TotalTime) &&
			stringsMatch(params.Yeild, recipe.Yield)

		// Ingredients
		for _, ingredient := range params.Ingredients {
			match = match && strings.Contains(recipe.Ingredients, ingredient)
		}

		if match {
			result = append(result, internal.SearchResult{RecipeName: recipe.RecipeName, RecipeId: id})
		}
	}

	sort.Slice(result, func(a, b int) bool {
		return result[a].RecipeName < result[b].RecipeName
	})

	return result
}
