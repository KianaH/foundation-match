package data

import (
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/pgxpool"
)

type Models struct {
	Foundation FoundationModel
}

func InitModels(dbpool *pgxpool.Pool) Models {
	return Models{
		Foundation: FoundationModel{DBPool: dbpool},
	}
}
