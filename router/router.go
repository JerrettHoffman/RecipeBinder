package router

import (
	"RecipeBinder/internal"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	router.Mux.HandleFunc("POST /create", router.createPostRecipeHandler)
	router.Mux.HandleFunc("GET /edit/{id}", router.editGetRecipeHandler)
	router.Mux.HandleFunc("POST /edit/{id}", router.editPostRecipeHandler)
	router.Mux.HandleFunc("GET /search", router.searchGetRecipeHandler)
	router.Mux.HandleFunc("GET /add", router.addGetRecipeHandler)
	readTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/read.tmpl"))
	editTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/edit.tmpl"))
	searchTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/search.tmpl"))
	addTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/add.tmpl"))
}

func parseID(r *http.Request) (internal.ID, error) {
	idStr := r.PathValue("id")

	if idStr != "" {
		if parsedId, err := strconv.Atoi(idStr); err == nil {
			return internal.ID(parsedId), nil
		} else {
			log.Printf("Failed to parse id %v\n", err)
			return -1, err
		}
	}

	return -1, errors.New("Request id was empty")
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

	id, err := parseID(r)
	if err != nil {
		http.Redirect(w, r, "/search", http.StatusNotFound)
		return
	}

	builder := internal.TestRecipeDataStrategy{}
	recipeData := builder.ReadRecipe(id)

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

type editTemplateData struct {
	SubmitURL   string
	RecipeName  string
	Author      string
	Uploader    string
	PrepTime    time.Duration
	TotalTime   time.Duration
	Yield       string
	Ingredients string
	Image       string
	Steps       string
}

func (router *Router) editGetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Redirect(w, r, "/search", http.StatusNotFound)
		return
	}

	builder := internal.TestRecipeDataStrategy{}
	recipeData := builder.ReadRecipe(id)

	data := editTemplateData{
		SubmitURL:   fmt.Sprintf("/edit/%d", id),
		RecipeName:  recipeData.RecipeName,
		Author:      recipeData.Author,
		Uploader:    recipeData.Uploader,
		PrepTime:    recipeData.PrepTime,
		TotalTime:   recipeData.TotalTime,
		Yield:       recipeData.Yield,
		Ingredients: recipeData.Ingredients,
		Image:       recipeData.Image,
		Steps:       recipeData.Steps,
	}

	if err := editTpl.Execute(w, data); err != nil {
		log.Printf("Failed to execute editGet %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

func fillDataFromForm(r *http.Request) (internal.RecipeData, error) {
	var err error

	if err = r.ParseForm(); err != nil {
		log.Printf("Failed to parse form: %v", err)
		return internal.RecipeData{}, err
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
		log.Printf("Failed to parse prepTime \"%s\": %s", prepTimeStr, err)
		return internal.RecipeData{}, err
	}

	totalTime := time.Second
	if totalTime, err = time.ParseDuration(totalTimeStr); err != nil {
		log.Printf("Failed to parse totalTime \"%s\": %s", totalTimeStr, err)
		return internal.RecipeData{}, err
	}

	data := internal.RecipeData{
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

	return data, nil
}

func (router *Router) editPostRecipeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Redirect(w, r, "/search", http.StatusNotFound)
		return
	}

	dbData, err := fillDataFromForm(r)
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Send to DB
	builder := internal.TestRecipeDataStrategy{}
	if err = builder.UpdateRecipe(dbData, id); err != nil {
		http.Error(w, "Could not update recipe", http.StatusInternalServerError)
		return
	}

	// Reroute to the new read page for the created index
	http.Redirect(w, r, fmt.Sprintf("/read/%d", id), http.StatusSeeOther)
}

func (router *Router) searchGetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = r.ParseForm(); err != nil {
		log.Printf("Failed to parse form: %v", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	type searchFormData struct {
		RecipeName      string
		Author          string
		Uploader        string
		PrepTime        string
		TotalTime       string
		Yield           string
		IngredientCount string
		Ingredients     []string
	}

	// Pull the data from the form
	formData := searchFormData{
		RecipeName:      r.FormValue("recipe-name"),
		Author:          r.FormValue("author"),
		Uploader:        r.FormValue("uploader"),
		PrepTime:        r.FormValue("prep-time"),
		TotalTime:       r.FormValue("total-time"),
		Yield:           r.FormValue("yield"),
		IngredientCount: r.FormValue("ingredient-count"),
	}

	// Parse any non-string fields
	prepTime := time.Second
	if formData.PrepTime != "" {
		if prepTime, err = time.ParseDuration(formData.PrepTime); err != nil {
			log.Printf("Failed to parse prepTime \"%s\": %s", formData.PrepTime, err)
			http.Error(w, "Invalid form data: Prep Time", http.StatusBadRequest)
			return
		}
	}

	totalTime := time.Second
	if formData.TotalTime != "" {
		if totalTime, err = time.ParseDuration(formData.TotalTime); err != nil {
			log.Printf("Failed to parse totalTime \"%s\": %s", formData.TotalTime, err)
			http.Error(w, "Invalid form data: Total Time", http.StatusBadRequest)
			return
		}
	}

	if formData.IngredientCount != "" {
		if ingredientCount, err := strconv.Atoi(formData.IngredientCount); err == nil {
			ingredientCount = max(0, min(ingredientCount, maxIngredientCapacity))
			formData.Ingredients = make([]string, 0, ingredientCount)
			for i := range ingredientCount {
				formData.Ingredients = append(formData.Ingredients, r.FormValue(fmt.Sprintf("ingredient-%d", i)))
			}
		} else {
			log.Printf("Failed to parse ingredient-count \"%s\": %s", formData.IngredientCount, err)
			http.Error(w, "Invalid form data: Ingredient Count", http.StatusBadRequest)
			return
		}
	}

	searchParams := internal.SearchParams{
		RecipeName:   formData.RecipeName,
		AuthorName:   formData.Author,
		UploaderName: formData.Uploader,
		PrepTime:     prepTime,
		TotalTime:    totalTime,
		Yeild:        formData.Yield,
		Ingredients:  formData.Ingredients,
	}

	searcher := internal.TestSearch{}
	searchResults := searcher.Search(searchParams)

	type data struct {
		FormData searchFormData
		Results  []internal.SearchResult
	}

	pageData := data{
		FormData: formData,
		Results:  searchResults,
	}

	if err := searchTpl.Execute(w, pageData); err != nil {
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
	data := editTemplateData{
		SubmitURL:   "/create",
		RecipeName:  "",
		Author:      "",
		Uploader:    "",
		PrepTime:    time.Hour + time.Minute,
		TotalTime:   time.Hour + time.Minute,
		Yield:       "",
		Ingredients: "",
		Image:       "",
		Steps:       "",
	}

	if err := editTpl.Execute(w, data); err != nil {
		log.Printf("Failed to execute createGet %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

func (router *Router) createPostRecipeHandler(w http.ResponseWriter, r *http.Request) {
	dbData, err := fillDataFromForm(r)
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Send to DB
	builder := internal.TestRecipeDataStrategy{}
	id, err := builder.CreateRecipe(dbData, internal.ID(1))
	if err != nil {
		http.Error(w, "Could not create recipe", http.StatusInternalServerError)
		return
	}

	// Reroute to the new read page for the created index
	http.Redirect(w, r, fmt.Sprintf("/read/%d", id), http.StatusSeeOther)
}
