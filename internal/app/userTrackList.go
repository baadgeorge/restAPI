package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetUsersTracks(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}
	tracklist, err := h.services.ITracklistService.GetUserTracklist(userId)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK,
		map[string]interface{}{
			"message":   "user track list",
			"tracklist": tracklist,
		})
	return
}

func (h *Handler) AddTrackToTrackList(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}

	trackId, err := h.GetTrackId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}

	err = h.services.ITracklistService.AddTrackToUserTrackList(userId, trackId)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("track %d added to user %d tracklist", trackId, userId),
	})
	return
}

func (h *Handler) DeleteTrackFromTrackList(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}

	trackId, err := h.GetTrackId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}
	err = h.services.ITracklistService.DeleteTrackFromUserTracklist(userId, trackId)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "track deleted",
	})
	return
}

func (h *Handler) DeleteTrackList(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}
	err = h.services.ITracklistService.DeleteUserTrackList(userId)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user tracklist deleted",
	})
	return
}
