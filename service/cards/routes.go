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
}

func (h *Handler) HandleGetCards(c echo.Context) error {
	cards, err := h.store.GetCards()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cards)
}
