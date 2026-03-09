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
	authorId, err := findAuthorByName(recipe.Author)

	// TODO: Returns an error if no author found, need to distinguish between this error and others, for now, attempt to insert if any error thrown here.
	// if err != nil {
	// 	return -1, err
	// }

	if authorId == -1 {
		println("Create new author Record")

		authorId, err = insertAuthor(dbAuthor{
			Id:   "",
			Name: recipe.Author,
		})
		if err != nil {
			return -1, err
		}
	}

	println("Author id found: " + fmt.Sprintf("%d", authorId))

	println("Adding new recipe")

	newRecipeId, err := insertRecipe(dbRecipe{
		Id:             "",
		Name:           recipe.RecipeName,
		AuthorId:       authorId,
		UploaderId:     userId,
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

func (d DbRecipeDataStrategy) ReadRecipe(recipeId internal.ID) (internal.RecipeData, error) {

	recipe, err := getRecipeById(recipeId)

	if err != nil {
		return internal.RecipeData{}, fmt.Errorf("Error retrieving recipe from database %W", err)
	}

	author, authorErr := getAuthorById(recipe.AuthorId)

	if authorErr != nil {
		return internal.RecipeData{}, fmt.Errorf("Error retrieving author details: %W", authorErr)
	}

	uploader, uploaderErr := getUsernameById(recipe.UploaderId)

	if uploaderErr != nil {
		return internal.RecipeData{}, fmt.Errorf("Error retrieving user details: %W", uploaderErr)
	}

	recipeData := internal.RecipeData{
		RecipeName:  recipe.Name,
		Author:      author.Name,
		Uploader:    uploader,
		PrepTime:    recipe.PrepTime,
		TotalTime:   recipe.TotalTime,
		Yield:       recipe.Yeild,
		Ingredients: recipe.IngredientText,
		Image:       "",
		Steps:       recipe.Steps,
	}

	return recipeData, nil
}

// TODO: Needs to validate that recipeID and userID match before performing delete
func (d DbRecipeDataStrategy) DeleteRecipe(id internal.ID, userId internal.ID) error {
	return nil
}

type DbUserAuthDataStrategy struct{}

// Returns UserAuthData or error if no user of userName can be found
func (d DbUserAuthDataStrategy) ReadAuthUser(userName string) (internal.UserAuthData, error) {
	user, err := findUserByUserName(userName)

	if err != nil {
		return internal.UserAuthData{}, err
	}

	return internal.UserAuthData{
		Id:             user.Id,
		UserName:       user.Username,
		HashedPassword: user.HashedPassword,
	}, nil
}

func (d DbUserAuthDataStrategy) CreateAuthUser(userName string, hashedPw string) error {
	_, err := insertUser(dbUserAuth{
		Id:             -1,
		Username:       userName,
		HashedPassword: hashedPw,
	})
	if err != nil {
		return fmt.Errorf("Error creating new auth user: %W", err)
	}

	return nil
}

// Thought needed here about forgotten password
func (d DbUserAuthDataStrategy) UpdateAuthUser(currUserId internal.ID, newUser internal.UserAuthData) error {
	return nil
}
