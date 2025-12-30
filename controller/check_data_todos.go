package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type checkRequest struct {
	Done bool `json:"done"`
}

func NewCheckTodosController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id/check", func(ctx echo.Context) error {
		id := ctx.Param("id")
		var request checkRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)
		var DoneInt int
		if request.Done {
			DoneInt = 1
		}
		_, err := db.Exec(
			"UPDATE todos SET done = ? WHERE id = ?",
			DoneInt,
			id,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.String(http.StatusOK, "OKAY BRO")
	})
}
