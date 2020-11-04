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
		INSERT  INTO proposals(user_id, offer_id)
        SELECT  :user_id, :offer_id
        WHERE   
            NOT EXISTS (
                SELECT FROM offers AS o
                WHERE  o.offer_id = :offer_id
                AND    o.user_id = :user_id
            )
        RETURNING proposal_id, user_id, offer_id, created_on;
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

	return nil, &model.Created{p.ProposalID, p.CreatedOn}
}