package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type ReqTrackId struct {
	TrackId int `json:"track_id" binding:"required"`
}

type ReqAuthor struct {
	Author string `json:"author" binding:"required"`
}

type ReqTitle struct {
	Title string `json:"title" binding:"required"`
}

type ReqAuthorAndTitle struct {
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

func (h *Handler) GetUserId(c *gin.Context) (int, error) {
	reqUserId, ok := c.Get("user_id")
	if !ok {
		err := errors.New("can't get user_id from context")
		logrus.Debug(err)
		return 0, err
	}
	return reqUserId.(int), nil
}

func (h *Handler) GetTrackId(c *gin.Context) (int, error) {
	var reqTrackId ReqTrackId
	err := c.BindJSON(&reqTrackId)
	if err != nil {
		logrus.Debug("can't bind json", reqTrackId)
		return 0, err
	}
	return reqTrackId.TrackId, nil
}

func (h *Handler) GetTitle(c *gin.Context) (string, error) {
	var reqTitle ReqTitle
	err := c.BindJSON(&reqTitle)
	if err != nil {
		logrus.Debug("can't bind json", reqTitle)
		return "", err
	}
	return reqTitle.Title, nil
}

func (h *Handler) GetAuthor(c *gin.Context) (string, error) {
	var reqAuthor ReqAuthor
	err := c.BindJSON(&reqAuthor)
	if err != nil {
		logrus.Debug("can't bind json", reqAuthor)
		return "", err
	}
	return reqAuthor.Author, nil
}

func (h *Handler) GetAuthorAndTitle(c *gin.Context) (string, string, error) {
	var reqAuthorAndTitle ReqAuthorAndTitle
	err := c.BindJSON(&reqAuthorAndTitle)
	if err != nil {
		logrus.Debug("can't bind json", reqAuthorAndTitle)
		return "", "", err
	}
	return reqAuthorAndTitle.Author, reqAuthorAndTitle.Title, nil
}

func (h *Handler) CheckToken(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		logrus.Debug("empty auth header")
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "empty auth header",
		})
		return
	}
	headerParse := strings.Split(header, " ")
	if len(headerParse) != 2 || headerParse[0] != "Bearer" || len(headerParse[1]) == 0 {
		logrus.Debug("wrong auth header")
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "wrong auth header",
		})
		return
	}
	userId, err := h.services.ITokenService.Parse(headerParse[1])
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't parse token",
		})
		return
	}
	c.Set("user_id", userId)
	return
}
