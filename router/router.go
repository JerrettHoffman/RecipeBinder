package router

import (
	"RecipeBinder/internal"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"text/template"
	"time"
)

const (
	maxIngredientCapacity = 100
	headerMarkup          = "## "
	bulletMarkup          = "* "
)

var (
	readTpl   *template.Template
	editTpl   *template.Template
	searchTpl *template.Template
	addTpl    *template.Template
)

type ingredientSection struct {
	Header      string
	Ingredients []string
}

type stepSection struct {
	Header string
	Steps  []string
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
	router.Mux.HandleFunc("GET /create", router.createGetRecipeHandler)
	router.Mux.HandleFunc("GET /edit/{id}", router.editGetRecipeHandler)
	router.Mux.HandleFunc("POST /edit/{id}", router.editPostRecipeHandler)
	router.Mux.HandleFunc("GET /search", router.searchGetRecipeHandler)
	router.Mux.HandleFunc("GET /add", router.addGetRecipeHandler)
	readTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/read.tmpl"))
	editTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/edit.tmpl"))
	searchTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/search.tmpl"))
	addTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/add.tmpl"))
}

// Take in string of ingredient text and separate it into sections with headers
// and individual ingredients
func formatIngredientSections(ingredientText string) []ingredientSection {
	outIngredientSections := make([]ingredientSection, 1, 2)

	sectionIndex := 0
	currentSection := &outIngredientSections[sectionIndex]
	currentSection.Ingredients = make([]string, 0, 3)

	lines := regexp.MustCompile("\r?\n").Split(ingredientText, -1)
	for _, line := range lines {
		if trimmedHeader, foundPrefix := strings.CutPrefix(line, headerMarkup); foundPrefix {
			outIngredientSections = append(outIngredientSections, ingredientSection{
				Header:      trimmedHeader,
				Ingredients: make([]string, 0, 3),
			})
			sectionIndex++
			currentSection = &outIngredientSections[sectionIndex]
		} else if trimmedBullet, foundPrefix := strings.CutPrefix(line, bulletMarkup); foundPrefix {
			currentSection.Ingredients = append(currentSection.Ingredients, trimmedBullet)
		}
	}

	return outIngredientSections
}

// Take in string of steps text and separate it into sections with headers
// and individual steps
func formatStepSections(stepText string) []stepSection {
	outStepSections := make([]stepSection, 1)

	sectionIndex := 0
	currentSection := &outStepSections[sectionIndex]
	currentSection.Steps = make([]string, 0, 3)

	lines := regexp.MustCompile("\r?\n").Split(stepText, -1)
	for _, line := range lines {
		if trimmedHeader, foundPrefix := strings.CutPrefix(line, headerMarkup); foundPrefix {
			outStepSections = append(outStepSections, stepSection{
				Header: trimmedHeader,
				Steps:  make([]string, 0, 3),
			})
			sectionIndex++
			currentSection = &outStepSections[sectionIndex]
		} else if len(line) > 0 {
			currentSection.Steps = append(currentSection.Steps, line)
		}
	}

	return outStepSections
}

func (router *Router) readRecipeHandler(w http.ResponseWriter, r *http.Request) {
	type data struct {
		RecipeName         string
		Author             string
		Uploader           string
		PrepTime           string
		TotalTime          string
		Yield              string
		IngredientSections []ingredientSection
		Image              string
		StepSections       []stepSection
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

	stepsSections := formatStepSections(recipeData.Steps)

	pageData := data{
		recipeData.RecipeName,
		recipeData.Author,
		recipeData.Uploader,
		prepTimeFormatted,
		totalTimeFormatted,
		recipeData.Yield,
		ingredientSections,
		recipeData.Image,
		stepsSections,
	}

	if err := readTpl.Execute(w, pageData); err != nil {
		log.Printf("Failed to execute read %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

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
	log.Printf("%v", dbData)

	// Reroute to the new read page for the created index
}

func (router *Router) searchGetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	if err := searchTpl.Execute(w, nil); err != nil {
		log.Printf("Failed to execute searchGet %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

func (router *Router) addGetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	if err := addTpl.Execute(w, nil); err != nil {
		log.Printf("Failed to execute addGet %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

func (router *Router) createGetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	if err := editTpl.Execute(w, nil); err != nil {
		log.Printf("Failed to execute createGet %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}
