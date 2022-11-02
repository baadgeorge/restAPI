package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"test/internal/datastruct"
)

type logIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) LogIn(c *gin.Context) {
	var input logIn
	err := c.BindJSON(&input)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't bing json",
		})
		return
	}
	token, err := h.services.IAuthService.Login(input.Email, input.Password)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "login error",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
	return
}

func (h *Handler) SignUp(c *gin.Context) {
	var input datastruct.FullUser

	err := c.BindJSON(&input)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't bing json",
		})
		return
	}
	id, err := h.services.IAuthService.SignUp(input)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "signUp error",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": id,
	})
	return
}
