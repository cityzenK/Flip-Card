package api

import (
	"cards/service/cards"
	"database/sql"
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

func (s *APIServer) Run() error {
	router := echo.New()
	defer router.Close()
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	subrouter := router.Group("/api/v1")
	CardsStore := cards.NewCards(s.db)
	cardsHandler := cards.NewHandle(CardsStore)
	cardsHandler.RegisterRouter(subrouter)

	return http.ListenAndServe(s.addr, router)
}
