package handlers

import (
	"net/http"

	"github.com/aleyrizvi/onoff-api/cmd/dto"
	"github.com/labstack/echo/v4"
)

type AnagramCheckReader interface {
	IsAnagram(word string, anagram string) bool
}

type AnagramHandler struct {
	Anagram AnagramCheckReader
}

func NewAnagramHandler(anagramService AnagramCheckReader) *AnagramHandler {
	return &AnagramHandler{
		Anagram: anagramService,
	}
}

func (a *AnagramHandler) HandleCheckAnagram(c echo.Context) error {
	var req dto.CheckAnagramRequestBody

	// bind request body to struct
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "A valid word and anagram is required",
		})
	}

	// validate body
	if err := c.Validate(req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.CheckAnagramResponseBody{
		IsAnagram: a.Anagram.IsAnagram(req.Word, req.Anagram),
	})
}
