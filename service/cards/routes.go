package cards

import (
	"cards/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store types.CardsInterface
}

func NewHandle(cards types.CardsInterface) *Handler {
	return &Handler{store: cards}
}

func (h *Handler) RegisterRouter(router *echo.Group) {
	router.GET("/cards", h.HandleGetCards)
	router.GET("/filters-cards/:filter-name", h.HandlerFilterCards)
}

func (h *Handler) HandleGetCards(c echo.Context) error {
	cards, err := h.store.GetCards()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cards)
}

func (h *Handler) HandlerFilterCards(c echo.Context) error {
	filters := map[string]string{
		"all":     "where 1 = 1",
		"general": "where type = 1",
		"code":    "where type = 2",
		"known":   "where known = 1",
		"unknown": "where known = 0",
	}

	statement := filters[c.Param("filter-name")]
	cards, err := h.store.GetCardsFilter(statement)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cards)
}
