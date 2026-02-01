package internal

import "time"

type RecipeData struct {
	RecipeName  string
	Author      string
	Uploader    string
	PrepTime    time.Duration
	TotalTime   time.Duration
	Yield       string
	Ingredients string
	Image       string
	Steps       string
}

type RecipdeDataBuilder interface {
	BuildRecipe() RecipeData
}
