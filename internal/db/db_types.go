package db

type dbAuthor struct {
	Id   string
	Name string
}

type dbUser struct {
	Id             string
	Username       string
	HashedPassword string
}

type dbRecipe struct {
	Id             string
	Name           string
	AuthorId       int `db: "author_id"`
	UploaderId     int `db: "uploader_id"`
	PrepTime       int
	TotalTime      int
	Steps          string
	IngredientText string
	Yeild          string
}
