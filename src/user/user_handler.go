package user

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang-echo-showcase/src/user/response"
	sqlc "golang-echo-showcase/src/user/sqlc/output"
)

type Handler struct {
	Service *Service
}

type ServiceInterface interface {
	CreateUser(ctx context.Context, firstname, lastname string) (sqlc.User, error)
	GetUser(ctx context.Context, id int64) (sqlc.User, error)
	UpdateUserFirstname(ctx context.Context, id int64, firstname string) error
	UpdateUserLastname(ctx context.Context, id int64, lastname string) error
	UpdateUser(ctx context.Context, id int64, firstname, lastname string) error
	DeleteUser(ctx context.Context, id int64) error
}

type UserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type UpdateUserRequest struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

type UpdateUserFirstnameRequest struct {
	Firstname string `json:"firstname"`
}

type UpdateUserLastnameRequest struct {
	Lastname string `json:"lastname"`
}

// ---------- Response Helper ----------

func successResponse(data interface{}, message string, start time.Time) *response.ApiResponse {
	return &response.ApiResponse{
		Success:        true,
		ProcessingTime: time.Since(start).String(),
		Data:           data,
		Message:        message,
	}
}

func errorResponse(message string, start time.Time) *response.ApiResponse {
	return &response.ApiResponse{
		Success:        false,
		ProcessingTime: time.Since(start).String(),
		Error:          message,
	}
}

// ---------- Handler ----------

func (h *Handler) SaveUser(c echo.Context) error {
	start := time.Now()

	firstname := c.QueryParam("firstname")
	lastname := c.QueryParam("lastname")

	if firstname == "" || lastname == "" {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse("Parameter 'firstname' und 'lastname' sind erforderlich.", start),
		)
	}

	user, err := h.Service.CreateUser(c.Request().Context(), firstname, lastname)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			errorResponse(err.Error(), start),
		)
	}

	return c.JSON(
		http.StatusCreated,
		successResponse(user, "Benutzer erfolgreich erstellt", start),
	)
}

func (h *Handler) GetUser(c echo.Context) error {
	start := time.Now()

	idStr := c.QueryParam("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse("Ungültige ID", start),
		)
	}

	user, err := h.Service.GetUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			errorResponse("Benutzer nicht gefunden", start),
		)
	}

	return c.JSON(
		http.StatusOK,
		successResponse(user, "Benutzer erfolgreich abgerufen", start),
	)
}

func (h *Handler) UpdateUserFirstname(c echo.Context) error {
	start := time.Now()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse("Ungültige ID", start),
		)
	}

	firstname := c.QueryParam("firstname")
	if firstname == "" {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse("Parameter 'firstname' ist erforderlich", start),
		)
	}

	err = h.Service.UpdateUserFirstname(c.Request().Context(), id, firstname)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			errorResponse("Fehler beim Aktualisieren des Benutzers", start),
		)
	}

	return c.JSON(
		http.StatusOK,
		successResponse(nil, "Vorname erfolgreich aktualisiert", start),
	)
}

func (h *Handler) UpdateUserLastname(c echo.Context) error {
	start := time.Now()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse("Ungültige ID", start),
		)
	}

	var req UpdateUserLastnameRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse("Ungültige Anfrage", start),
		)
	}

	if req.Lastname == "" {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse("Nachname ist erforderlich", start),
		)
	}

	err = h.Service.UpdateUserLastname(c.Request().Context(), id, req.Lastname)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			errorResponse("Fehler beim Aktualisieren des Benutzers", start),
		)
	}

	return c.JSON(
		http.StatusOK,
		successResponse(nil, "Nachname erfolgreich aktualisiert", start),
	)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	start := time.Now()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse("Ungültige ID", start),
		)
	}

	err = h.Service.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			errorResponse("Fehler beim Löschen des Benutzers", start),
		)
	}

	return c.JSON(
		http.StatusOK,
		successResponse(nil, "Benutzer erfolgreich gelöscht", start),
	)
}

func (h *Handler) ListUsers(c echo.Context) error {
	start := time.Now()

	return c.JSON(
		http.StatusNotImplemented,
		errorResponse("Noch nicht implementiert", start),
	)
}
