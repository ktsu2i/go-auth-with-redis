package handler

import (
	"backend/model"
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

	return c.JSON(http.StatusOK, u)
}
