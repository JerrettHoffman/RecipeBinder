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

type TestRecipeBuilder struct {
}

func (t TestRecipeBuilder) BuildRecipe() RecipeData {
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
		Ingredients: "4 Cups lorem, A pinch of ipsum",
		Image:       "lorem ipsum image",
		Steps:       "1. Lorem \n 2.Ipsum \n 3.I hope the new lines work?",
	}
}
