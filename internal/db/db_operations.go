package db

import (
	"RecipeBinder/internal"
	"fmt"
	"strings"
)

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

func insertUser(user dbUserAuth) (internal.ID, error) {
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

func updateSingleRecipeValues(recipe dbRecipe) error {
	q := dbQuery{
		query: `
		UPDATE recipes
		SET name = @recipeName, author_id = @authorId, uploader_id = @uploaderId, prep_time = @prepTime, total_time = @totalTime, steps = @steps, ingredient_text = @ingredientText, yeild = @yeild
		WHERE id=@id`,
		args: dbInsertArgs{
			"recipeName":     recipe.Name,
			"authorId":       recipe.AuthorId,
			"uploaderId":     recipe.UploaderId,
			"prepTime":       recipe.PrepTime,
			"totalTime":      recipe.TotalTime,
			"steps":          recipe.Steps,
			"ingredientText": recipe.IngredientText,
			"yeild":          recipe.Yeild,
			"id":             recipe.Id,
		},
	}

	err := q.dbExec()
	if err != nil {
		return err
	}

	return nil
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
		SELECT id, name, author_id, uploader_id, prep_time, total_time, steps, ingredient_text, yeild FROM recipes
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

func getUserById(userId internal.ID) (dbUserAuth, error) {
	q := dbQuery{
		query: `
		SELECT id, username FROM users
		WHERE id=@id`,
		args: dbInsertArgs{
			"id": userId,
		},
	}
	user, err := q.dbQueryReturningSingleAuthUser()
	if err != nil {
		return dbUserAuth{}, err
	}

	return user, nil
}

func findUserByUserName(userName string) (dbUserAuth, error) {
	q := dbQuery{
		query: `
		SELECT * FROM users
		WHERE username=@username`,
		args: dbInsertArgs{
			"username": userName,
		},
	}
	user, err := q.dbQueryReturningSingleAuthUser()
	if err != nil {
		return dbUserAuth{}, fmt.Errorf("Error finding user by username: %W", err)
	}
	return user, nil
}

func getUsernameById(userId internal.ID) (string, error) {
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
		return "", err
	}

	return user.Username, nil
}

func constructSearchSQL(params internal.SearchParams) dbQuery {
	args := dbInsertArgs{}
	prefix := `SELECT recipes.id, recipes.name FROM recipes `
	var whereClauses []string
	var joinClauses []string

	if params.AuthorName != `` {
		joinClauses = append(joinClauses, `INNER JOIN authors ON recipes.author_id=authors.id`)
		whereClauses = append(whereClauses, `authors.name >= @authorName`)
		args["authorName"] = params.AuthorName
	}

	if params.UploaderName != `` {
		joinClauses = append(joinClauses, `INNER JOIN users ON recipes.uploader_id=users.id`)
		whereClauses = append(whereClauses, `users.username >= @userName`)
		args["userName"] = params.UploaderName
	}

	// assemble where clauses and concatinate to where string
	// TODO CLOVE: what will you pass if the user doesn't specify values to search by
	if params.PrepTime > 0 {
		whereClauses = append(whereClauses, `recipes.prep_time = @prepTime`)
		args["prepTime"] = params.PrepTime
	}

	if params.TotalTime > 0 {
		whereClauses = append(whereClauses, `recipes.total_time = @totalTime`)
		args["totalTime"] = params.TotalTime
	}

	if params.Yeild != `` {
		whereClauses = append(whereClauses, `recipes.yeild = @yeild`)
		args["yeild"] = params.Yeild
	}

	if params.RecipeName != `` {
		whereClauses = append(whereClauses, `recipes.name >= @recipeName`)
		args["recipeName"] = params.RecipeName
	}
	// TODO JERRETT: figure out text search in here too

	// put everything together
	var queryString strings.Builder
	queryString.WriteString(prefix)

	for index, value := range joinClauses {
		if index > 0 {
			queryString.WriteString(" ")
		}
		queryString.WriteString(value)
	}

	//add a space between join and where clauses
	queryString.WriteString(" ")

	for index, value := range whereClauses {
		if index == 0 {
			queryString.WriteString("WHERE ")
		}
		if index > 0 {
			queryString.WriteString(" AND ")
		}
		queryString.WriteString(value)
	}

	return dbQuery{
		query: queryString.String(),
		args:  args,
	}

}

func searchQueryReturningMultipleIds(params internal.SearchParams) ([]internal.SearchResult, error) {
	q := constructSearchSQL(params)

	dbRes, err := q.dbQueryMultipleRowsReturningIds()
	if err != nil {
		return nil, err
	}

	return dbRes, nil

}
