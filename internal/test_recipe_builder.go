package internal

import "time"

// Implements RecipdeDataStrategy
type TestRecipeDataStrategy struct {
}

func (t TestRecipeDataStrategy) ReadRecipe(id ID) RecipeData {
	prepTime, err := time.ParseDuration("1h30m")
	if err != nil {
		prepTime = 0
	}

	totalTime, err := time.ParseDuration("2h45m")
	if err != nil {
		totalTime = 0
	}

	return RecipeData{
		RecipeName:  "Lorem Ipsum",
		Author:      "Mr. Lorem",
		Uploader:    "Ms. Ipsum",
		PrepTime:    prepTime,
		TotalTime:   totalTime,
		Yield:       "4 Dolor",
		Ingredients: "* 4 Cups lorem\r\n* A pinch of ipsum",
		Image:       "lorem ipsum image",
		Steps:       "1. Lorem \n 2.Ipsum \n 3.I hope the new lines work?",
	}
}

func (t TestRecipeDataStrategy) UpdateRecipe(recipe RecipeData, recipeId ID) error {
	return nil
}

func (t TestRecipeDataStrategy) CreateRecipe(recipe RecipeData) (ID, error) {
	return 0, nil
}

func (t TestRecipeDataStrategy) DeleteRecipe(recipeID ID) error {
	return nil
}
