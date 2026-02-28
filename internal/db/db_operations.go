package db

import "RecipeBinder/internal"

func insertAuthor(author DbAuthor) (internal.ID, error) {
	q := dbQuery{
		query: `
		INSERT INTO authors (name)
		VALUES (@name)
		RETURNING id`,
		args: dbInsertArgs{
			"name": author.Name,
		},
	}

	id, err := q.dbQuerySingleRowReturningId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func insertUser(user DbUser) (internal.ID, error) {
	q := dbQuery{
		query: `
		INSERT INTO users (username, hashed_password)
		VALUES (@userName, @hashedPassword)
		RETURNING id`,
		args: dbInsertArgs{
			"userName":       user.Username,
			"hashedPassword": user.HashedPassword,
		},
	}
	id, err := q.dbQuerySingleRowReturningId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func insertRecipe(recipe DbRecipe) (internal.ID, error) {
	q := dbQuery{
		query: `
		INSERT INTO recipes (name, author_id, uploader_id, prep_time, total_time, steps, ingredient_text, yeild)
		VALUES (@recipeName, @authorId, @uploaderId, @prepTime, @totalTime, @steps, @ingredientText, @yeild)
		RETURNING id`,
		args: dbInsertArgs{
			"recipeName":     recipe.Name,
			"authorId":       recipe.AuthorId,
			"uploaderId":     recipe.UploaderId,
			"prepTime":       recipe.PrepTime,
			"totalTime":      recipe.TotalTime,
			"steps":          recipe.Steps,
			"ingredientText": recipe.IngredientText,
			"yeild":          recipe.Yeild,
		},
	}
	id, err := q.dbQuerySingleRowReturningId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func insertIngredient(ingredient DbIngredient) (internal.ID, error) {
	q := dbQuery{
		query: `
		INSERT INTO ingredients (name)
		VALUES (@ingredientName)
		RETURNING id`,
		args: dbInsertArgs{
			"ingredientName": ingredient.Name,
		},
	}
	id, err := q.dbQuerySingleRowReturningId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func insertRecipeIngredient(recipeIngredient DbRecipeIngredient) error {
	q := dbQuery{
		query: `
		INSERT INTO recipe_ingredients (recipe_id, ingredient_id)
		VALUES (@recipeId, @ingredientId)`,
		args: dbInsertArgs{
			"recipeId":     recipeIngredient.RecipeId,
			"ingredientId": recipeIngredient.IngredientId,
		},
	}
	err := q.dbExec()

	if err != nil {
		return err
	}

	return nil
}

func findAuthor(authorName string) (internal.ID, error) {
	q := dbQuery{
		query: `
		SELECT id FROM authors
		WHERE name=@name`,
		args: dbInsertArgs{
			"name": authorName,
		},
	}

	id, err := q.dbQuerySingleRowReturningId()
	if err != nil {
		return -1, nil
	}

	return id, nil
}
