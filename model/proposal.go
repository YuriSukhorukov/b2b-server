package model

type AddProposal struct {
	OfferID 	string 		`db:"offer_id" json:"offer_id" example:"1d586b05-7b80-4a3a-bf2c-ce48169d4e85" format:"string"`
}

func (a AddProposal) Validation() error {
	switch {
	case len(a.OfferID) == 0:
		return ErrOfferIDInvalid
	default:
		return nil
	}
}

type Proposal struct {
	ProposalID 	string 		`db:"proposal_id" json:"proposal_id" example:"1d586b05-7b80-4a3a-bf2c-ce48169d4e85" format:"string"`
	UserID 		string    	`db:"user_id" json:"user_id" example:"1d586b05-7b80-4a3a-bf2c-ce48169d4e85" format:"string"`
	OfferID 	string 		`db:"offer_id" json:"offer_id" example:"1d586b05-7b80-4a3a-bf2c-ce48169d4e85" format:"string"`
	CreatedOn 	string 		`db:"created_on" json:"created_on" example:"2020-10-28T14:58:56.059Z" format:"string"`
}