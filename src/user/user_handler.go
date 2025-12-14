package user

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"golang-echo-showcase/src/user/sqlc/output"
	"context"
	"golang-echo-showcase/src/user/response"
)

type Handler struct {
	Queries *sqlc.Queries 
}

type UserRequest struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
}

func successResponse(data interface{}, message string) *response.ApiResponse {
    return &response.ApiResponse{
        Success: true,
        Data:    data,
        Message: message,
    }
}

func errorResponse(message string) *response.ApiResponse {
    return &response.ApiResponse{
        Success: false,
        Error:   message,
    }
}

func (h* Handler) SaveUser(c echo.Context) error {
	var req UserRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Ungültig")
	}

	params := sqlc.CreateUserParams{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
	}

	user, err := h.Queries.CreateUser(context.Background(), params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Fehler beim Speichern.")
	}

	return c.JSON(http.StatusCreated, user)
}

func (h* Handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusCreated, id+" wurde aktualisiert (aus Handler-Paket)")
}

func (h* Handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusCreated, id+" wurde aktualisiert (aus Handler-Paket)")
}

func  (h* Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Benutzer mit ID: "+id+" gelöscht (aus Handler-Paket)")
}