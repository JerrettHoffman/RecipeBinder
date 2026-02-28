package db

import "RecipeBinder/internal"

type DbRecipeDataStrategy struct{}

// We assume that the username in RecipeData uploader field comes from an already created user
func (d DbRecipeDataStrategy) CreateRecipe(recipe internal.RecipeData, userId internal.ID) (internal.ID, error) {
	// order matters since some tables use IDs from earlier tables as foriegn keys

	//check if author exists, otherwise create new one
	_, err := insertAuthor(DbAuthor{
		Id:   "",
		Name: recipe.Author,
	})
	if err != nil {
		return -1, err
	}

	//insert Recipe with userids provided and author id from previous step

	return 0, nil
}

// TODO: Needs to validate that recipe ID and User ID match before performing update
func (d DbRecipeDataStrategy) UpdateRecipe(recipe internal.RecipeData, recipeId internal.ID, userId internal.ID) error {
	// what to do if update orphans an author record?
	return nil
}

func (d DbRecipeDataStrategy) ReadRecipe(id internal.ID) (internal.RecipeData, error) {

	return internal.RecipeData{}, nil
}

// TODO: Needs to validate that recipeID and userID match before performing delete
func (d DbRecipeDataStrategy) DeleteRecipe(id internal.ID, userId internal.ID) error {
	return nil
}

type DbUserAuthDataStrategy struct{}

func (d DbUserAuthDataStrategy) ReadAuthUser(userName string) (internal.UserAuthData, error) {
	return internal.UserAuthData{}, nil
}
func (d DbUserAuthDataStrategy) CreateAuthUser(userName string, hashedPw string) error {
	return nil
}
func (d DbUserAuthDataStrategy) UpdateAuthUser(currUserId internal.ID, newUser internal.UserAuthData) error {
	return nil
}
