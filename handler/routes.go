package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SmokierLemur51/minecraft-wms/data"
)

// Load all sales routes
func (h *CoreHandler) LoadRoutes(router *http.ServeMux) *http.ServeMux {
	// Error page
	router.HandleFunc("GET /error/{error}", h.ErrorHandler)

	// Html page get requests
	router.HandleFunc("GET /", h.IndexPageHandler)
	router.HandleFunc("GET /worlds", h.WorldsPagehandler)
	router.HandleFunc("GET /worlds/{world}", h.WorldHandler)

	// Post requests
	router.HandleFunc("POST /worlds/create-new-world", h.NewWorldHandler)
	router.HandleFunc("POST /worlds/search", h.SearchWorlds)

	// Testing Routes
	router.HandleFunc("GET /test/create-tables", h.CreateTables)
	router.HandleFunc("GET /test/insert", h.TestInsert)

	return router
}

// Error
func (h *CoreHandler) ErrorHandler(w http.ResponseWriter, r *http.Request) {
	page := HtmlTemplate{Page: "error.html", Error: r.PathValue("error")}
	page.RenderTemplate(w)
}

// Testing
func (h *CoreHandler) CreateTables(w http.ResponseWriter, r *http.Request) {
	data.CreateSchema(h.DB)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *CoreHandler) TestInsert(w http.ResponseWriter, r *http.Request) {
	data.PopulateWorlds(h.DB)
	data.PopulateLocationCategories(h.DB)
	http.Redirect(w, r, "/worlds", http.StatusSeeOther)
}

// Handlers

func (h *CoreHandler) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	// Render index page
	index := HtmlTemplate{Page: "index.html", Title: "Logan Lee"}
	index.RenderTemplate(w)
}

func (h *CoreHandler) WorldsPagehandler(w http.ResponseWriter, r *http.Request) {

	p := HtmlTemplate{
		Page:  "worlds.html",
		Title: "Your Worlds",
	}
	p.RenderTemplate(w)

}

func (h *CoreHandler) NewWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Process the world form for errors
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	newWorld := r.FormValue("world")
	fmt.Println(newWorld)
	http.Redirect(w, r, "/worlds", http.StatusSeeOther)
}

func (h *CoreHandler) WorldHandler(w http.ResponseWriter, r *http.Request) {
	var world data.World
	var err error
	err = r.ParseForm()
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error/%s", err), http.StatusBadRequest)
	}
	if world, err = data.LoadWorldByName(h.DB, r.FormValue("world")); err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error/%s", err), http.StatusNotFound)
	} else {
		p := HtmlTemplate{
			Page:  "worlds.html",
			Title: "Your Worlds",
			World: world,
		}
		p.RenderTemplate(w)
	}
}

// Redirect you to the world, doesn't care if it exists.
func (h *CoreHandler) SearchWorlds(w http.ResponseWriter, r *http.Request) {
	search := r.PathValue("search")
	http.Redirect(w, r, "/worlds/"+search, http.StatusSeeOther)
}
