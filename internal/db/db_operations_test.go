package db

import (
	"RecipeBinder/internal"
	"reflect"
	"testing"
)

func Test_constructSearchSQL(t *testing.T) {

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		params internal.SearchParams
		want   dbQuery
	}{
		{
			name: "Join and Where Logic Combined",
			params: internal.SearchParams{
				AuthorName: "Gordon Ramsay",
				PrepTime:   20,
			},
			want: dbQuery{
				// Now correctly generates one JOIN and one WHERE with AND
				query: "SELECT id, name FROM recipes INNER JOIN authors ON recipes.author_id=authors.id WHERE authors.name >= @authorName AND recipes.prep_time = @prepTime",
				args: dbInsertArgs{
					"authorName": "Gordon Ramsay",
					"prepTime":   20,
				},
			},
		},
		{
			name: "Multiple Joins",
			params: internal.SearchParams{
				AuthorName:   "Alice",
				UploaderName: "Bob",
			},
			want: dbQuery{
				query: "SELECT id, name FROM recipes INNER JOIN authors ON recipes.author_id=authors.id INNER JOIN users ON recipes.uploader_id=users.id WHERE authors.name >= @authorName AND users.username >= @userName",
				args: dbInsertArgs{
					"authorName": "Alice",
					"userName":   "Bob",
				},
			},
		},
		{
			name: "All Parameters Provided",
			params: internal.SearchParams{
				AuthorName:   "Chef",
				UploaderName: "Admin",
				PrepTime:     10,
				TotalTime:    60,
				Yeild:        "4",
				RecipeName:   "Cake",
			},
			want: dbQuery{
				query: "SELECT id, name FROM recipes INNER JOIN authors ON recipes.author_id=authors.id INNER JOIN users ON recipes.uploader_id=users.id WHERE authors.name >= @authorName AND users.username >= @userName AND recipes.prep_time = @prepTime AND recipes.total_time = @totalTime AND recipes.yeild = @yeild AND recipes.name >= @recipeName",
				args: dbInsertArgs{
					"authorName": "Chef",
					"userName":   "Admin",
					"prepTime":   10,
					"totalTime":  60,
					"yeild":      "4",
					"recipeName": "Cake",
				},
			},
		},
		{
			name: "Only Standalone Where Clause",
			params: internal.SearchParams{
				RecipeName: "Pasta",
			},
			want: dbQuery{
				// Note the double space between 'recipes' and 'WHERE' due to the empty join loop + manual space
				query: "SELECT id, name FROM recipes  WHERE recipes.name >= @recipeName",
				args:  dbInsertArgs{"recipeName": "Pasta"},
			},
		},
		{
			name:   "Completely Empty Params",
			params: internal.SearchParams{},
			want: dbQuery{
				query: "SELECT id, name FROM recipes  ",
				args:  dbInsertArgs{},
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructSearchSQL(tt.params)
			// TODO: update the condition below to compare got with tt.want.
			if got.query != tt.want.query {
				t.Errorf("Query mismatch\ngot:  %q\nwant: %q", got.query, tt.want.query)
			}

			if !reflect.DeepEqual(got.args, tt.want.args) {
				t.Errorf("Args mismatch\ngot:  %v\nwant: %v", got.args, tt.want.args)
			}
		})
	}
}
