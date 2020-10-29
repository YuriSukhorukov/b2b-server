package model

import "errors"

type AddOffer struct {
	Title   			string    	`db:"title" json:"title" example:"Сгущенка" format:"string"`
	Description 		string 		`db:"description" json:"description" example:"Оригинальная сгущенка Рогачев" format:"string"`
	Price				string 		`db:"price" json:"price" example:"1000000" format:"int"`
	Amount 				string 		`db:"amount" json:"amount" example:"100" format:"int"`
	CurrencyCode 		string 		`db:"currency_code" json:"currency_code" example:"RUB" format:"string"`
	OfferType 			string 		`db:"offer_type" json:"offer_type" example:"BUY" format:"string"`
	MeasureUnitCode 	string 		`db:"measure_unit_code" json:"measure_unit_code" example:"KG" format:"string"`
	DateExpires 		string 		`db:"date_expires" json:"date_expires" example:"2020-10-28T14:58:56.059Z" format:"string"`
	Country 			string 		`db:"country" json:"country" example:"Российская Федерация" format:"string"`
	City 				string 		`db:"city" json:"city" example:"Москва" format:"string"`
}

var (
	ErrTitleInvalid = errors.New("title is empty")
	ErrDescriptionInvalid = errors.New("description is empty")
	ErrPriceInvalid = errors.New("price is empty")
	ErrAmountInvalid = errors.New("amount is empty")
	ErrCurrencyCodeInvalid = errors.New("currencyCode is empty")
	ErrOfferTypeInvalid = errors.New("offerType is empty")
	ErrMeasureUnitCodeInvalid = errors.New("measureUnitCode is empty")
	ErrDateExpiresInvalid = errors.New("dateExpires is empty")
	ErrCountryInvalid = errors.New("country is empty")
	ErrCityInvalid = errors.New("city is empty")
)

func (a AddOffer) Validation() error {
	switch {
	case len(a.Title) == 0:
		return ErrTitleInvalid
	case len(a.Description) == 0:
		return ErrDescriptionInvalid
	case len(a.Price) == 0:
		return ErrPriceInvalid
	case len(a.Amount) == 0:
		return ErrAmountInvalid
	case len(a.CurrencyCode) == 0:
		return ErrCurrencyCodeInvalid
	case len(a.OfferType) == 0:
		return ErrOfferTypeInvalid
	case len(a.MeasureUnitCode) == 0:
		return ErrMeasureUnitCodeInvalid
	case len(a.DateExpires) == 0:
		return ErrDateExpiresInvalid
	case len(a.Country) == 0:
		return ErrCountryInvalid
	case len(a.City) == 0:
		return ErrCityInvalid
	default:
		return nil
	}
}

type Offer struct {
	OfferID   			string    	`db:"offer_id" json:"offer_id" example:"1d586b05-7b80-4a3a-bf2c-ce48169d4e85" format:"string"`
	UserID   			string    	`db:"user_id" json:"user_id" example:"1d586b05-7b80-4a3a-bf2c-ce48169d4e85" format:"string"`
	Title   			string    	`db:"title" json:"title" example:"Сгущенка" format:"string"`
	Description 		string 		`db:"description" json:"description" example:"Оригинальная сгущенка Рогачев" format:"string"`
	Price				string 		`db:"price" json:"price" example:"1000000" format:"int"`
	Amount 				string 		`db:"amount" json:"amount" example:"100" format:"int"`
	CurrencyCode 		string 		`db:"currency_code" json:"currency_code" example:"RUB" format:"string"`
	OfferType 			string 		`db:"offer_type" json:"offer_type" example:"BUY" format:"string"`
	MeasureUnitCode 	string 		`db:"measure_unit_code" json:"measure_unit_code" example:"KG" format:"string"`
	DateExpires 		string 		`db:"date_expires" json:"date_expires" example:"2020-10-28T14:58:56.059Z" format:"string"`
	Country 			string 		`db:"country" json:"country" example:"Российская Федерация" format:"string"`
	City 				string 		`db:"city" json:"city" example:"Москва" format:"string"`
	CreatedOn			string 		`db:"created_on" json:"created_on" example:"2020-10-28T14:58:56.059Z" format:"string"`
}