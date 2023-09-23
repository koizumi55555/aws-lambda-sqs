package master_repo

import (
	db "aws-lambda-sqs/db/postgres"
)

type MasterRepository struct {
	DBHandler *db.DBHandler
}

func New(dbh *db.DBHandler) *MasterRepository {
	return &MasterRepository{DBHandler: dbh}
}
