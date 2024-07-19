package api

import (
	"cards/service/cards"
	"database/sql"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func (s *APIServer) Run() error {
	router := echo.New()
	defer router.Close()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("cmd/web/*.html")),
	}
	router.Renderer = renderer
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	subrouter := router.Group("/api/v1")
	CardsStore := cards.NewCards(s.db)
	cardsHandler := cards.NewHandle(CardsStore)
	cardsHandler.RegisterRouter(subrouter)

	return http.ListenAndServe(s.addr, router)
}
