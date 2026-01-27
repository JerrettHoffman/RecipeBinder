package router

import (
	"RecipeBinder/internal"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
	"time"
)

const (
	maxIngredientCapacity = 100
)

type Router struct {
	Mux *http.ServeMux
}

func (router *Router) Setup() {
	router.Mux = http.NewServeMux()

	// Set up routing
	fs := http.FileServer(http.Dir("assets"))
	router.Mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router.Mux.HandleFunc("/read/{id}", router.readRecipeHandler)
	router.Mux.HandleFunc("GET /edit/{id}", router.editGetRecipeHandler)
	router.Mux.HandleFunc("POST /edit/{id}", router.editPostRecipeHandler)
}

var readTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/read.tmpl"))

func (router *Router) readRecipeHandler(w http.ResponseWriter, r *http.Request) {
	type data struct {
		RecipeName  string
		Author      string
		Uploader    string
		PrepTime    string
		TotalTime   string
		Yield       int
		Ingredients []string
		Image       string
		Steps       []string
	}

	// Todo: use id for actual lookup of data
	builder := internal.TestRecipeBuilder{}
	recipeData := builder.BuildRecipe()

	prepTimeHours := int(recipeData.PrepTime.Hours())
	prepTimeMinutes := int(recipeData.PrepTime.Minutes()) % 60
	totalTimeHours := int(recipeData.TotalTime.Hours())
	totalTimeMinutes := int(recipeData.TotalTime.Minutes()) % 60

	pageData := data{
		recipeData.RecipeName,
		recipeData.Author,
		recipeData.Uploader,
		fmt.Sprintf("%dh %dm", prepTimeHours, prepTimeMinutes),
		fmt.Sprintf("%dh %dm", totalTimeHours, totalTimeMinutes),
		recipeData.Yield,
		recipeData.Ingredients,
		recipeData.Image,
		nil,
	}

	pageData.Steps = regexp.MustCompile("\r?\n").Split(recipeData.Steps, -1)

	if err := readTpl.Execute(w, pageData); err != nil {
		log.Printf("Failed to execute read %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

var editTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/edit.tmpl"))

func (router *Router) editGetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// Todo: use id for actual lookup of data
	builder := internal.TestRecipeBuilder{}
	recipeData := builder.BuildRecipe()

	if err := editTpl.Execute(w, recipeData); err != nil {
		log.Printf("Failed to execute editGet %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

func (router *Router) editPostRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = r.ParseForm(); err != nil {
		log.Printf("Failed to parse form in editPost: %v", err)
	}

	// Pull the data from the form
	recipeName := r.PostFormValue("recipe-name")
	author := r.PostFormValue("author")
	uploader := r.PostFormValue("uploader")
	prepTimeStr := r.PostFormValue("prep-time")
	totalTimeStr := r.PostFormValue("total-time")
	yieldStr := r.PostFormValue("yield")
	finalImage := r.PostFormValue("final-image")
	ingredientCountStr := r.PostFormValue("ingredient-count")
	steps := r.PostFormValue("steps")

	// Parse any non-string fields
	prepTime := time.Second
	if prepTime, err = time.ParseDuration(prepTimeStr); err != nil {
		log.Printf("Failed to parse prepTime \"%s\" in editPost: %s", prepTimeStr, err)
	}

	totalTime := time.Second
	if totalTime, err = time.ParseDuration(totalTimeStr); err != nil {
		log.Printf("Failed to parse totalTime \"%s\" in editPost: %s", totalTimeStr, err)
	}

	yield := 0
	if yield, err = strconv.Atoi(yieldStr); err != nil {
		log.Printf("Failed to parse yield \"%s\" in editPost: %s", yieldStr, err)
	}

	dbData := internal.RecipeData{
		RecipeName:  recipeName,
		Author:      author,
		Uploader:    uploader,
		PrepTime:    prepTime,
		TotalTime:   totalTime,
		Yield:       yield,
		Ingredients: nil,
		Image:       finalImage,
		Steps:       steps,
	}

	// Parse the ingredients one by one
	if ingredientCount, err := strconv.Atoi(ingredientCountStr); err != nil {
		log.Printf("Failed to parse ingredient count \"%s\" in editPost: %v", ingredientCountStr, err)
	} else {
		// Reserve capacity for the number of expected ingredients
		dbData.Ingredients = make([]string, 0, min(ingredientCount, maxIngredientCapacity))
		for i := range ingredientCount {
			val := r.PostFormValue(fmt.Sprintf("ingredient-%d", i))
			dbData.Ingredients = append(dbData.Ingredients, val)
		}
	}

	// Send to DB

	// Reroute to the new read page for the created index
}
