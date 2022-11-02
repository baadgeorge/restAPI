package repository

import "github.com/jackc/pgx"

type Repository struct {
	IUserQuery
	ITrackQuery
	ITracklistQuery
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		IUserQuery:      NewUserQuery(db),
		ITrackQuery:     NewTrackQuery(db),
		ITracklistQuery: NewTracklistQuery(db),
	}
}
