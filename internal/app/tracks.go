package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"test/internal/datastruct"
)

func (h *Handler) CreateTrack(c *gin.Context) {
	var input datastruct.Track
	if err := c.BindJSON(&input); err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can't bind json",
		})
		return
	}
	trackId, err := h.services.ITrackService.CreateTrack(&input)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't create track",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"track_id": trackId,
	})
	return
}

func (h *Handler) GetTrackById(c *gin.Context) {
	trackId, err := h.GetTrackId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't get track_id",
		})
		return
	}

	track, err := h.services.ITrackService.GetTrackById(trackId)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can't get track by id",
		})
		return
	}
	c.JSON(http.StatusOK, track)
}

func (h *Handler) GetTracksByAuthorAndTitle(c *gin.Context) {
	author, title, err := h.GetAuthorAndTitle(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't get author and title",
		})
		return
	}

	tracks, err := h.services.ITrackService.GetTracksByAuthorAndTitle(author, title)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can't get tracks by author and title",
		})
		return
	}
	c.JSON(http.StatusOK, tracks)
}

func (h *Handler) GetTracksByAuthor(c *gin.Context) {
	author, err := h.GetAuthor(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't get author",
		})
		return
	}

	tracks, err := h.services.ITrackService.GetTracksByAuthor(author)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can't get tracks by author",
		})
		return
	}
	c.JSON(http.StatusOK, tracks)
}

func (h *Handler) GetTracksByTitle(c *gin.Context) {
	title, err := h.GetTitle(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't get title",
		})
		return
	}

	tracks, err := h.services.ITrackService.GetTracksByTitle(title)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can't get tracks by title",
		})
		return
	}
	c.JSON(http.StatusOK, tracks)
}

func (h *Handler) UpdateTrack(c *gin.Context) {
	var input datastruct.Track
	if err := c.BindJSON(&input); err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can't bind json",
		})
		return
	}
	if err := h.services.ITrackService.UpdateTrack(input); err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can't update tracks info",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "track updated",
	})
}

func (h *Handler) DeleteTrack(c *gin.Context) {
	trackId, err := h.GetTrackId(c)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't get track_id",
		})
		return
	}
	err = h.services.ITrackService.DeleteTrack(trackId)
	if err != nil {
		logrus.Debug(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can't delete track",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "track deleted",
	})
}
