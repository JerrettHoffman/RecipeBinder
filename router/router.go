package router

import (
	"RecipeBinder/internal"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"
	"time"
)

const (
	maxIngredientCapacity = 100
	headerMarkup = "##"
	bulletMarkup = "* "
)

type ingredientSection struct {
	Header      string
	Ingredients []string
}

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

// Take in string of ingredient text and separate it into sections with headers
// and individual ingredients
func formatIngredientSections(ingredientText string) []ingredientSection {
	return nil
}

func (router *Router) readRecipeHandler(w http.ResponseWriter, r *http.Request) {
	type data struct {
		RecipeName          string
		Author              string
		Uploader            string
		PrepTime            string
		TotalTime           string
		Yield               string
		ingredientsSections []ingredientSection
		Image               string
		Steps               []string
	}

	// Todo: use id for actual lookup of data
	builder := internal.TestRecipeBuilder{}
	recipeData := builder.BuildRecipe()

	// Format the times
	prepTimeHours := int(recipeData.PrepTime.Hours())
	prepTimeMinutes := int(recipeData.PrepTime.Minutes()) % 60
	prepTimeFormatted := fmt.Sprintf("%dh %dm", prepTimeHours, prepTimeMinutes)

	totalTimeHours := int(recipeData.TotalTime.Hours())
	totalTimeMinutes := int(recipeData.TotalTime.Minutes()) % 60
	totalTimeFormatted := fmt.Sprintf("%dh %dm", totalTimeHours, totalTimeMinutes)

	ingredientSections := formatIngredientSections(recipeData.Ingredients)

	steps := regexp.MustCompile("\r?\n").Split(recipeData.Steps, -1)

	pageData := data{
		recipeData.RecipeName,
		recipeData.Author,
		recipeData.Uploader,
		prepTimeFormatted,
		totalTimeFormatted,
		recipeData.Yield,
		ingredientSections,
		recipeData.Image,
		steps,
	}

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
	yield := r.PostFormValue("yield")
	finalImage := r.PostFormValue("final-image")
	ingredientsStr := r.PostFormValue("ingredient")
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

	dbData := internal.RecipeData{
		RecipeName:  recipeName,
		Author:      author,
		Uploader:    uploader,
		PrepTime:    prepTime,
		TotalTime:   totalTime,
		Yield:       yield,
		Ingredients: ingredientsStr,
		Image:       finalImage,
		Steps:       steps,
	}

	// Send to DB
	log.Printf("%v", dbData);

	// Reroute to the new read page for the created index
}
