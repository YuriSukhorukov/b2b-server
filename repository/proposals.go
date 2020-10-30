package repository

import (
	"github.com/jmoiron/sqlx"
	// "github.com/b2b-server/model"
)

type ProposalRepository struct {
	db *sqlx.DB
}

func NewProposalRepository(db *sqlx.DB) *ProposalRepository {
	return &ProposalRepository{db}
}