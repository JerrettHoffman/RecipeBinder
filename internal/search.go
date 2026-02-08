package internal

import "time"

// Assume we are "and-ing" all of the search parameters
type SearchParams struct {
	recipeName   string
	authorName   string
	uploaderName string
	prepTime     time.Duration
	totalTime    time.Duration
	yeild        string
	ingredients  []string
}

type SearchResult struct {
	recipeName string
	recipeId   ID
}

type RecipeSearch interface {
	search(params SearchParams) []SearchResult
}

type TestSearch struct{}

func (t TestSearch) search(params SearchParams) []SearchResult {
	return []SearchResult{{
		recipeName: "Test Recipe 1",
		recipeId:   1,
	}, {
		recipeName: "Test Recipe 2",
		recipeId:   2,
	}}
}
