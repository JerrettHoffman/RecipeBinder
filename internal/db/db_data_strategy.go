package internal

import "RecipeBinder/internal"

type DbRecipeDataStrategy struct{}

func (d DbRecipeDataStrategy) CreateRecipe(recipe internal.RecipeData) (internal.ID, error) {

	return 0, nil
}

func (d DbRecipeDataStrategy) UpdateRecipe(recipe internal.RecipeData, recipeId internal.ID) error {
	return nil
}

func (d DbRecipeDataStrategy) ReadRecipe(id internal.ID) (internal.RecipeData, error) {

	return internal.RecipeData{}, nil
}

func (d DbRecipeDataStrategy) DeleteRecipe(id internal.ID) error {
	return nil
}
