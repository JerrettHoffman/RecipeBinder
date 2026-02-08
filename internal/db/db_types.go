package internal

type DbIngredient struct {
	Id   string
	Name string
}

type DbAuthor struct {
	Id   string
	Name string
}

type DbUser struct {
	Id          string
	Username    string
	DisplayName string
}

type DbRecipe struct {
	Id             string
	Name           string
	AuthorId       int
	UploaderId     int
	PrepTime       int
	TotalTime      int
	Steps          string
	IngredientText string
	Yeild          string
}

type DbRecipeIngredient struct {
	RecipeId     string
	IngredientId string
}
