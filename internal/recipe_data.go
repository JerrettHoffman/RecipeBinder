package internal

type ID = int

type RecipeData struct {
	RecipeName  string
	Author      string
	Uploader    string
	PrepTime    int // Duration of preptime in minutes
	TotalTime   int // Duration of TotalTime in minutes
	Yield       string
	Ingredients string
	Image       string
	Steps       string
}

type RecipeDataStrategy interface {
	ReadRecipe(recipeId ID) (RecipeData, error)
	UpdateRecipe(recipe RecipeData, recipeId ID, userId ID) error
	CreateRecipe(recipe RecipeData, uploaderId ID) (ID, error)
	DeleteRecipe(recipeId ID, userId ID) error
}
