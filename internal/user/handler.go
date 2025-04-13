package user

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

type UserHandler struct {
    service UserService
}

func NewUserHandler(e *echo.Echo, service UserService) {
    handler := &UserHandler{service}

    e.POST("/register", handler.Register)
    e.POST("/login", handler.Login)
}

type RegisterRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *UserHandler) Register(c echo.Context) error {
    var req RegisterRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
    }

    user, err := h.service.Register(req.Name, req.Email, req.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c echo.Context) error {
    var req LoginRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
    }

    user, err := h.service.Login(req.Email, req.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
    }

    return c.JSON(http.StatusOK, user)
}

