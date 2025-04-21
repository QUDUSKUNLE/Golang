package repo

import (

	"github.com/QUDUSKUNLE/microservices/hospital-service/adapters/db"
)

type Repository struct {
	database *db.Queries
}
