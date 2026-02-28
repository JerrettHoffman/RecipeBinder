package db

import (
	"RecipeBinder/internal"
	"fmt"
)

type DbRecipeDataStrategy struct{}

// We assume that the username in RecipeData uploader field comes from an already created user
func (d DbRecipeDataStrategy) CreateRecipe(recipe internal.RecipeData, userId internal.ID) (internal.ID, error) {
	// order matters since some tables use IDs from earlier tables as foriegn keys

	//check if author exists, otherwise create new one
	println("Check if author exists")
	authorId, err := findAuthor(recipe.Author)
	if err != nil {
		return -1, err
	}

	if authorId == -1 {
		println("Create new author Record")

		authorId, err = insertAuthor(DbAuthor{
			Id:   "",
			Name: recipe.Author,
		})
		if err != nil {
			return -1, err
		}
	}

	println("Author id found: " + fmt.Sprintf("%d", authorId))

	println("Adding new recipe")

	newRecipeId, err := insertRecipe(DbRecipe{
		Id:             "",
		Name:           recipe.RecipeName,
		AuthorId:       int(authorId),
		UploaderId:     int(userId),
		PrepTime:       int(recipe.PrepTime),
		TotalTime:      int(recipe.TotalTime),
		Steps:          recipe.Steps,
		IngredientText: recipe.Ingredients,
		Yeild:          recipe.Yield,
	})

	if err != nil {
		return -1, err
	}

	return newRecipeId, nil

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
