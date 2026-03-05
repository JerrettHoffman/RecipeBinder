package router

import (
	"RecipeBinder/internal"
	"RecipeBinder/internal/auth"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags

	envErr := godotenv.Load("../../.env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v", envErr)
	} else {
		m.Run()
	}
}

func Test_parseID(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r         *http.Request
		pathValue string
		want      internal.ID
		wantErr   bool
	}{
		{
			"Parse id from read endpoint",
			httptest.NewRequest("GET", "/read", nil),
			"1",
			1,
			false,
		},
		{
			"Bad id to read endpoint",
			httptest.NewRequest("GET", "/read", nil),
			"bad",
			0,
			true,
		},
		{
			"Parse id from edit endpoint",
			httptest.NewRequest("GET", "/edit", nil),
			"2",
			2,
			false,
		},
		{
			"Bad id to edit endpoint",
			httptest.NewRequest("GET", "/edit", nil),
			"fail",
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.SetPathValue("id", tt.pathValue)
			got, gotErr := parseID(tt.r)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("parseID() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("parseID() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("parseID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasValidSession(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r        *http.Request
		userData auth.UserData
		addData  bool
		want     bool
	}{
		{
			"Valid user data",
			httptest.NewRequest("GET", "/read", nil),
			auth.UserData{Id: 1, User: "user"},
			true,
			true,
		},
		{
			"Invalid user id",
			httptest.NewRequest("GET", "/read", nil),
			auth.UserData{Id: auth.UninitialzedId, User: ""},
			true,
			false,
		},
		{
			"No data in context",
			httptest.NewRequest("GET", "/read", nil),
			auth.UserData{Id: auth.UninitialzedId, User: ""},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.addData {
				ctx := context.WithValue(context.Background(), userDataContext, tt.userData)
				tt.r = tt.r.WithContext(ctx)
			}
			got := hasValidSession(tt.r)
			if got != tt.want {
				t.Errorf("hasValidSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			stepText: "## Section one\r\nTest this out\r\nWith another line\r\n",
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
			stepText: "## Section one\r\nTest this out\r\nWith another line\r\n## Section two\r\nNew lines added\r\n",
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
			stepText: "First line\r\n## Section one\r\nTest this out\r\nWith another line\r\n",
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
			stepText: "First line\r\n\r\n\r\n## Section one\r\nTest this out\r\nWith another line\r\n",
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
			ingredientText: "## Section One\r\n* 1 Tablespoon salt\r\n",
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
			ingredientText: "## Section One\r\n* 1 Tablespoon salt\r\n## Section Two\r\n* 2 Cups flour\r\n",
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

func Test_formatDuration(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		duration time.Duration
		want     string
	}{
		{"Test hours only", time.Hour, "1h"},
		{"Test minutes only", time.Minute, "1m"},
		{"Test hours and minutes", time.Hour + time.Minute, "1h 1m"},
		{"Test ignore seconds", time.Minute + time.Second, "1m"},
		{"Test negative hours", -time.Hour, ""},
		{"Test negative minutes", -time.Minute, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatDuration(tt.duration)
			if got != tt.want {
				t.Errorf("formatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_readRecipeHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w              *httptest.ResponseRecorder
		r              *http.Request
		pathValue      string
		wantStatusCode int
	}{
		{
			"Bad id request",
			httptest.NewRecorder(),
			httptest.NewRequest("GET", "/read/id", nil),
			"id",
			http.StatusFound,
		},
		{
			"Valid request (only checking no error)",
			httptest.NewRecorder(),
			httptest.NewRequest("GET", "/read/1", nil),
			"1",
			http.StatusOK,
		},
	}

	router := Router{}
	router.Setup()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.SetPathValue("id", tt.pathValue)
			router.readRecipeHandler(tt.w, tt.r)
			if tt.w.Code != tt.wantStatusCode {
				t.Errorf("formatIngredientSections() = %v, want %v", tt.w.Code, tt.wantStatusCode)
			}
		})
	}
}

func TestRouter_editGetRecipeHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w              *httptest.ResponseRecorder
		ctx            context.Context
		method         string
		target         string
		pathValue      string
		wantStatusCode int
	}{
		{
			"Bad id request",
			httptest.NewRecorder(),
			context.WithValue(context.Background(), userDataContext, auth.UserData{Id: auth.UninitialzedId, User: "user"}),
			"GET",
			"/edit/id",
			"id",
			http.StatusFound,
		},
		{
			"Bad user session",
			httptest.NewRecorder(),
			context.WithValue(context.Background(), userDataContext, auth.UserData{Id: auth.UninitialzedId, User: "user"}),
			"GET",
			"/edit/1",
			"1",
			http.StatusFound,
		},
		{
			"Missing user session",
			httptest.NewRecorder(),
			context.Background(),
			"GET",
			"/edit/1",
			"1",
			http.StatusFound,
		},
		{
			"Unauthorized edit",
			httptest.NewRecorder(),
			context.WithValue(context.Background(), userDataContext, auth.UserData{Id: 1, User: "user"}),
			"GET",
			"/edit/1",
			"1",
			http.StatusFound,
		},
		{
			"Valid request (only checking no error)",
			httptest.NewRecorder(),
			context.WithValue(context.Background(), userDataContext, auth.UserData{Id: 1, User: "Ms. Ipsum"}),
			"GET",
			"/edit/1",
			"1",
			http.StatusOK,
		},
	}

	router := Router{}
	router.Setup()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequestWithContext(tt.ctx, tt.method, tt.target, nil)
			req.SetPathValue("id", tt.pathValue)
			router.editGetRecipeHandler(tt.w, req)
			if tt.w.Code != tt.wantStatusCode {
				t.Errorf("formatIngredientSections() = %v, want %v", tt.w.Code, tt.wantStatusCode)
			}
		})
	}
}

func Test_fillDataFromForm(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		recipeName     string
		author         string
		prepTimeStr    string
		totalTimeStr   string
		yield          string
		finalImage     string
		ingredientsStr string
		steps          string
		want           internal.RecipeData
		wantErr        bool
	}{
		{
			"Valid form",
			"Recipe",
			"Author",
			"1m",
			"1m",
			"Yield",
			"FinalImage",
			"* Ingredients",
			"Steps",
			internal.RecipeData{
				RecipeName:  "Recipe",
				Author:      "Author",
				Uploader:    "",
				PrepTime:    time.Minute,
				TotalTime:   time.Minute,
				Yield:       "Yield",
				Ingredients: "* Ingredients",
				Image:       "FinalImage",
				Steps:       "Steps",
			},
			false,
		},
		{
			"Bad prep time",
			"Recipe",
			"Author",
			"bad",
			"1m",
			"Yield",
			"FinalImage",
			"* Ingredients",
			"Steps",
			internal.RecipeData{
				RecipeName:  "Recipe",
				Author:      "Author",
				Uploader:    "",
				PrepTime:    time.Minute,
				TotalTime:   time.Minute,
				Yield:       "Yield",
				Ingredients: "* Ingredients",
				Image:       "FinalImage",
				Steps:       "Steps",
			},
			true,
		},
		{
			"Bad total time",
			"Recipe",
			"Author",
			"1m",
			"bad",
			"Yield",
			"FinalImage",
			"* Ingredients",
			"Steps",
			internal.RecipeData{
				RecipeName:  "Recipe",
				Author:      "Author",
				Uploader:    "",
				PrepTime:    time.Minute,
				TotalTime:   time.Minute,
				Yield:       "Yield",
				Ingredients: "* Ingredients",
				Image:       "FinalImage",
				Steps:       "Steps",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formDataStr := fmt.Sprintf(
				"recipe-name=%s&author=%s&prep-time=%s&total-time=%s&yield=%s&final-image=%s&ingredient=%s&steps=%s",
				tt.recipeName,
				tt.author,
				tt.prepTimeStr,
				tt.totalTimeStr,
				tt.yield,
				tt.finalImage,
				tt.ingredientsStr,
				tt.steps,
			)
			print(formDataStr)
			body := strings.NewReader(formDataStr)
			req := httptest.NewRequest("POST", "/edit", body)
			req.ContentLength = int64(len(formDataStr))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			got, gotErr := fillDataFromForm(req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("fillDataFromForm() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("fillDataFromForm() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("fillDataFromForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: Need test for edit POST, search GET
