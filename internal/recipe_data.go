package internal

import "time"

type recipeData struct {
	recipeName  string
	author      string
	uploader    string
	prepTime    time.Duration
	totalTime   time.Duration
	yeild       int
	ingredients []string
	image       string
	steps       string
}

type RecipdeDataBuilder interface {
	BuildRecipe() recipeData
}

type testRecipeBuilder struct {
}

func (t testRecipeBuilder) BuildRecipe() recipeData {
	prepTime, err := time.ParseDuration("1h30m")
	if err != nil {
		prepTime = 0
	}

	totalTime, err := time.ParseDuration("2h45m")
	if err != nil {
		totalTime = 0
	}

	return recipeData{
		recipeName:  "Lorem Ipsum",
		author:      "Mr. Lorem",
		uploader:    "Ms. Ipsum",
		prepTime:    prepTime,
		totalTime:   totalTime,
		yeild:       4,
		ingredients: []string{"lorem", "ipsum", "turkey neck"},
		image:       "lorem ipsum image",
		steps:       "1. Lorem /n 2.Ipsum /n 3.I hope the new lines work?",
	}
}
