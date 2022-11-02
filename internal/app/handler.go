package app

import (
	"github.com/gin-gonic/gin"
	"test/internal/service"
)

type Handler struct {
	services *service.DataService
}

func NewHandler(services *service.DataService) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.LogIn)
		auth.POST("/sign-up", h.SignUp)
	}
	api := router.Group("/api", h.CheckToken)
	{
		tracks := api.Group("/tracks")
		{
			tracks.POST("/create_track", h.CreateTrack)
			tracks.GET("/get_track_by_id", h.GetTrackById)
			tracks.GET("/get_tracks_by_author_and_title", h.GetTracksByAuthorAndTitle)
			tracks.GET("/get_tracks_by_author", h.GetTracksByAuthor)
			tracks.GET("/get_tracks_by_title", h.GetTracksByTitle)
			tracks.PUT("/update_track", h.UpdateTrack)
			tracks.DELETE("/delete_track", h.DeleteTrack)
		}
		users := api.Group("/users")
		{
			person := users.Group("/person")
			{
				person.GET("/get_user", h.GetUserById)
				person.GET("/get_email_by_user", h.GetEmailByUserId)
				person.PUT("/update_user", h.UpdateUser)
				person.DELETE("/delete_user", h.DeleteUser)
			}
			trackList := users.Group("/tracklist")
			{
				trackList.GET("/get_tracklist", h.GetUsersTracks)
				trackList.POST("/add_track_to_tracklist", h.AddTrackToTrackList)
				trackList.DELETE("/delete_track_from_tracklist", h.DeleteTrackFromTrackList)
				trackList.DELETE("/delete_tracklist", h.DeleteTrackList)
			}
		}
	}
	return router
}
