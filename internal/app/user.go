package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"test/internal/datastruct"
)

func (h *Handler) GetUserById(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}
	user, err := h.services.IUserService.GetUser(userId)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var input datastruct.User
	if err := c.BindJSON(&input); err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}

	err := h.services.IUserService.UpdateUser(input)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
	return
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}
	err = h.services.IUserService.DeleteUser(userId)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user deleted",
	})
	return
}

func (h *Handler) GetEmailByUserId(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}
	email, err := h.services.IUserService.GetEmailByUserId(userId)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userId,
		"email":   email,
	})
	return
}
