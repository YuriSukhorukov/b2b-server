package model

type AddProposal struct {
	Title   			string    	`db:"title" json:"title" example:"Сгущенка" format:"string"`
}

func (a AddProposal) Validation() error {
	switch {
	case len(a.Title) == 0:
		return ErrTitleInvalid
	default:
		return nil
	}
}

type Proposal struct {
	OfferID   			string    	`db:"offer_id" json:"offer_id" example:"1d586b05-7b80-4a3a-bf2c-ce48169d4e85" format:"string"`
}