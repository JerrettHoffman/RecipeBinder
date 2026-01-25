package router

import (
	"RecipeBinder/internal"
	"log"
	"net/http"
	"text/template"
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
}

var readTpl = template.Must(template.ParseFiles("templates/base.tmpl", "templates/header.tmpl", "templates/read.tmpl"))

func (router *Router) readRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// Todo: use id for actual lookup of data 
	builder := internal.TestRecipeBuilder{}
	recipeData := builder.BuildRecipe()

	if err := readTpl.Execute(w, recipeData); err != nil {
		log.Printf("Failed to execute read %v\n", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}
