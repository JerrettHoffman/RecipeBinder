package internal

import "time"

type ID int

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

type RecipeDataStrategy interface {
	ReadRecipe(recipeId ID) (RecipeData, error)
	UpdateRecipe(recipe RecipeData, recipeId ID) error
	CreateRecipe(recipe RecipeData) (ID, error)
	DeleteRecipe(recipeId ID) error
}
