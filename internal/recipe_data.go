package internal

import "time"

type recipeData struct {
	RecipeName  string
	Author      string
	Uploader    string
	PrepTime    time.Duration
	TotalTime   time.Duration
	Yield       int
	Ingredients []ingredient
	Image       string
	Steps       string
}

type ingredient struct {
	IngredientName string
	Amount         string
}

type RecipdeDataBuilder interface {
	BuildRecipe() recipeData
}

type TestRecipeBuilder struct {
}

func (t TestRecipeBuilder) BuildRecipe() recipeData {
	prepTime, err := time.ParseDuration("1h30m")
	if err != nil {
		prepTime = 0
	}

	totalTime, err := time.ParseDuration("2h45m")
	if err != nil {
		totalTime = 0
	}

	return recipeData{
		RecipeName:  "Lorem Ipsum",
		Author:      "Mr. Lorem",
		Uploader:    "Ms. Ipsum",
		PrepTime:    prepTime,
		TotalTime:   totalTime,
		Yield:       4,
		Ingredients: BuildTestIngredients(),
		Image:       "lorem ipsum image",
		Steps:       "1. Lorem \n 2.Ipsum \n 3.I hope the new lines work?",
	}
}

func BuildTestIngredients() []ingredient {
	ingredient1 := ingredient{
		IngredientName: "Lorem",
		Amount:         "2.5 Lor",
	}
	ingredient2 := ingredient{
		IngredientName: "Ipsum",
		Amount:         "1/2 Pinch of Ip",
	}
	return []ingredient{ingredient1, ingredient2}
}
