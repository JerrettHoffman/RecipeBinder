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

func search(params SearchParams) []SearchResult {
	return []SearchResult{}
}
