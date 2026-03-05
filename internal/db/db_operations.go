package db

import "RecipeBinder/internal"

func insertAuthor(author dbAuthor) (internal.ID, error) {
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

func insertUser(user dbUser) (internal.ID, error) {
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

func insertRecipe(recipe dbRecipe) (internal.ID, error) {
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

func findAuthorByName(authorName string) (internal.ID, error) {
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
		return -1, err
	}

	return id, nil
}

func getRecipeById(recipeId internal.ID) (dbRecipe, error) {
	q := dbQuery{
		query: `
		SELECT * FROM recipes
		WHERE id=@id`,
		args: dbInsertArgs{
			"id": recipeId,
		},
	}

	recipe, err := q.dbQueryReturningSingleRecipe()
	if err != nil {
		return dbRecipe{}, err
	}

	return recipe, nil
}

func getAuthorById(authorId internal.ID) (dbAuthor, error) {
	q := dbQuery{
		query: `
		SELECT * FROM authors
		WHERE id=@id`,
		args: dbInsertArgs{
			"id": authorId,
		},
	}

	author, err := q.dbQueryReturningSingleAuthor()
	if err != nil {
		return dbAuthor{}, err
	}

	return author, nil
}

func getUserNameById(userId internal.ID) (dbUser, error) {
	q := dbQuery{
		query: `
		SELECT id, username FROM users
		WHERE id=@id`,
		args: dbInsertArgs{
			"id": userId,
		},
	}
	user, err := q.dbQueryReturningSingleUser()
	if err != nil {
		return dbUser{}, err
	}

	return user, nil
}
