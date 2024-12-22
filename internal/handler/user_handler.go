package handler

import (
	"errors"
	"net/http"
	"strconv"
	"user_service/internal/domain"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(userUseCase domain.UserService) *UserHandler {
	return &UserHandler{userService: userUseCase}
}

func (h *UserHandler) Register(e *echo.Echo) {
	e.POST("/users", h.CreateUser)
	e.GET("/users/:id", h.GetUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.POST("/auth", h.Authenticate)
}

func (h *UserHandler) Authenticate(c echo.Context) error {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	user, err := h.userService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	// Здесь вы можете создать JWT токен или сессию для аутентифицированного пользователя
	// Например:
	// token, err := createJWTToken(user)
	// if err != nil {
	//     return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create token"})
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Authentication successful",
		"user":    user,
		// "token":   token,
	})
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.userService.Create(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	user.ID = id

	if err := h.userService.Update(user); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	return c.JSON(http.StatusOK, user)
}
