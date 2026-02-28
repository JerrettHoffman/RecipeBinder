package db

import (
	i "RecipeBinder/internal"
	db "RecipeBinder/internal/db"
)

func DbTest() error {
	// Currently assume there is already a user created with the username "testuser"
	// userstrategy := db.DbUserAuthDataStrategy{}
	// userErr := userstrategy.CreateAuthUser("testuser")
	userId := i.ID(1)
	r1 := i.RecipeData{
		RecipeName:  "Test Recipe 1",
		Author:      "Author One",
		Uploader:    "testuser",
		PrepTime:    5000,
		TotalTime:   10000,
		Yield:       "2 tests",
		Ingredients: "Here we test the ingredients list: 1 Test , 2 Tests, 1/2 Test",
		Image:       "blah",
		Steps:       "First we test, then we beat the tests together with the tests, then gently fold in the remaing tests and but the test in the oven for 5 mins",
	}
	r2 := i.RecipeData{
		RecipeName:  "Test Recipe 2",
		Author:      "Author One",
		Uploader:    "testuser",
		PrepTime:    5000,
		TotalTime:   10000,
		Yield:       "4 tests",
		Ingredients: "Here we test the ingredients list: 1 Test , 2 Tests, 1/2 Test",
		Image:       "blah",
		Steps:       "First we test, then we beat the tests together with the tests, then gently fold in the remaing tests and but the test in the oven for 5 mins",
	}
	r3 := i.RecipeData{
		RecipeName:  "Test Recipe 3",
		Author:      "Author Two",
		Uploader:    "testuser",
		PrepTime:    5000,
		TotalTime:   10000,
		Yield:       "8 tests",
		Ingredients: "Here we test the ingredients list: 1 Test , 2 Tests, 1/2 Test",
		Image:       "blah",
		Steps:       "First we test, then we beat the tests together with the tests, then gently fold in the remaing tests and but the test in the oven for 5 mins",
	}

	strategy := db.DbRecipeDataStrategy{}
	_, err := strategy.CreateRecipe(r1, userId)
	if err != nil {
		return err
	}
	_, err = strategy.CreateRecipe(r2, userId)
	if err != nil {
		return err
	}
	_, err = strategy.CreateRecipe(r3, userId)
	if err != nil {
		return err
	}

	return nil
}
