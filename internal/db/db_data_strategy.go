package internal

import "RecipeBinder/internal"

type DbRecipeDataStrategy struct{}

// We assume that the username in RecipeData uploader field comes from an already created user
func (d DbRecipeDataStrategy) CreateRecipe(recipe internal.RecipeData) (internal.ID, error) {
	// order matters since some tables use IDs from earlier tables as foriegn keys
	userId, err := readUserId(recipe.Uploader)
	if err != nil {
		return -1, err
	}

	authorId, err := insertAuthor(DbAuthor{
		Id:   "",
		Name: recipe.Author,
	})
	if err != nil {
		return -1, err
	}

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
