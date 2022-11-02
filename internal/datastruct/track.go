package datastruct

type Track struct {
	ID       int    `json:"track_id" db:"track_id"`
	Author   string `json:"author" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Genre    string `json:"genre" binding:"required"`
	Album    string `json:"album" binding:"required"`
	Duration int    `json:"duration" binding:"required"`
}
