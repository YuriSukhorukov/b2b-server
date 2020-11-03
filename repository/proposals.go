package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/b2b-server/model"
)

type ProposalRepository struct {
	db *sqlx.DB
}

func NewProposalRepository(db *sqlx.DB) *ProposalRepository {
	return &ProposalRepository{db}
}

func (r ProposalRepository) InsertProposal(proposal model.Proposal) (error, *model.Created) {
	p := model.Proposal{}
	s := `
        
    `
    rows, err := r.db.NamedQuery(s, proposal)

    if err != nil {
    	return ErrSomethingWrong, nil
    }
	for rows.Next() {
   		err = rows.StructScan(&p)
	}
	if err != nil {
    	return ErrSomethingWrong, nil
    }

	return nil, &model.Created{p.OfferID, p.CreatedOn}
}