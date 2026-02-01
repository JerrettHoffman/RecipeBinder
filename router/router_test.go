package router

import (
	"slices"
	"testing"
)

func Test_formatStepSections(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		stepText string
		want     []stepSection
	}{
		{
			name:     "Single line",
			stepText: "Test this out\r\n",
			want: []stepSection{
				{
					Header: "",
					Steps:  []string{"Test this out"},
				},
			},
		},
		{
			name:     "Two line",
			stepText: "Test this out\r\nWith another line\r\n",
			want: []stepSection{
				{
					Header: "",
					Steps:  []string{"Test this out", "With another line"},
				},
			},
		},
		{
			name:     "One Section Start",
			stepText: "##Section one\r\nTest this out\r\nWith another line\r\n",
			want: []stepSection{
				{
					Header: "",
					Steps:  []string{},
				},
				{
					Header: "Section one",
					Steps:  []string{"Test this out", "With another line"},
				},
			},
		},
		{
			name:     "Two Sections Start",
			stepText: "##Section one\r\nTest this out\r\nWith another line\r\n##Section two\r\nNew lines added\r\n",
			want: []stepSection{
				{
					Header: "",
					Steps:  []string{},
				},
				{
					Header: "Section one",
					Steps:  []string{"Test this out", "With another line"},
				},
				{
					Header: "Section two",
					Steps:  []string{"New lines added"},
				},
			},
		},
		{
			name:     "Two Lines then One Section",
			stepText: "First line\r\n##Section one\r\nTest this out\r\nWith another line\r\n",
			want: []stepSection{
				{
					Header: "",
					Steps:  []string{"First line"},
				},
				{
					Header: "Section one",
					Steps:  []string{"Test this out", "With another line"},
				},
			},
		},
		{
			name:     "Extra newlines",
			stepText: "First line\r\n\r\n\r\n##Section one\r\nTest this out\r\nWith another line\r\n",
			want: []stepSection{
				{
					Header: "",
					Steps:  []string{"First line"},
				},
				{
					Header: "Section one",
					Steps:  []string{"Test this out", "With another line"},
				},
			},
		},
		{
			name:     "Empty",
			stepText: "",
			want: []stepSection{
				{
					Header: "",
					Steps:  []string{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatStepSections(tt.stepText)
			if !slices.EqualFunc(got, tt.want, func(gotSection, wantSection stepSection) bool {
				return gotSection.Header == wantSection.Header && slices.Equal(gotSection.Steps, wantSection.Steps)
			}) {
				t.Errorf("formatStepSections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatIngredientSections(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ingredientText string
		want           []ingredientSection
	}{
		{
			name:           "One ingredient",
			ingredientText: "* 1 Tablespoon salt\r\n",
			want: []ingredientSection{
				{
					Header:      "",
					Ingredients: []string{"1 Tablespoon salt"},
				},
			},
		},
		{
			name:           "Two ingredients",
			ingredientText: "* 1 Tablespoon salt\r\n* 2 Cups flour\r\n",
			want: []ingredientSection{
				{
					Header:      "",
					Ingredients: []string{"1 Tablespoon salt", "2 Cups flour"},
				},
			},
		},
		{
			name:           "One section",
			ingredientText: "##Section One\r\n* 1 Tablespoon salt\r\n",
			want: []ingredientSection{
				{
					Header:      "",
					Ingredients: []string{},
				},
				{
					Header:      "Section One",
					Ingredients: []string{"1 Tablespoon salt"},
				},
			},
		},
		{
			name:           "Two sections",
			ingredientText: "##Section One\r\n* 1 Tablespoon salt\r\n##Section Two\r\n* 2 Cups flour\r\n",
			want: []ingredientSection{
				{
					Header:      "",
					Ingredients: []string{},
				},
				{
					Header:      "Section One",
					Ingredients: []string{"1 Tablespoon salt"},
				},
				{
					Header:      "Section Two",
					Ingredients: []string{"2 Cups flour"},
				},
			},
		},
		{
			name:           "Extra newlines",
			ingredientText: "* 1 Tablespoon salt\r\n\r\n\r\n* 2 Cups flour\r\n",
			want: []ingredientSection{
				{
					Header:      "",
					Ingredients: []string{"1 Tablespoon salt", "2 Cups flour"},
				},
			},
		},
		{
			name:           "Malformed Lines",
			ingredientText: "asdf\r\n* 1 Tablespoon salt\r\n*2 Cups flour\r\naiojo\r\n",
			want: []ingredientSection{
				{
					Header:      "",
					Ingredients: []string{"1 Tablespoon salt"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatIngredientSections(tt.ingredientText)
			if !slices.EqualFunc(got, tt.want, func(gotSection, wantSection ingredientSection) bool {
				return gotSection.Header == wantSection.Header && slices.Equal(gotSection.Ingredients, wantSection.Ingredients)
			}) {
				t.Errorf("formatIngredientSections() = %v, want %v", got, tt.want)
			}
		})
	}
}
