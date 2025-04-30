package handler

import (
	"backend/model"
	"backend/pkg"
	"backend/repo"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c echo.Context) error {
	var req model.UserReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	UUID, err := uuid.NewRandom()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	u := model.User{
		ID:             UUID.String(),
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: string(hashed),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	r := repo.NewUserRepo(DB)
	if err := r.CreateUser(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Generate token pair
	pair, err := pkg.NewTokenPair(u.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Save refresh token to redis
	if err := repo.NewRedisRepo(RC).Save(c.Request().Context(), u.ID, pair.JTI, 7*24*60*60); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Set auth cookie
	SetAuthCookie(c.Response().Writer, pair)

	return c.JSON(http.StatusOK, u)
}
