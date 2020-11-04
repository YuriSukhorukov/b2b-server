package model

import "errors"

var (
	ErrEmailInvalid 				= errors.New("email is empty")
	ErrPasswordInvalid 				= errors.New("password is empty")
	ErrTitleInvalid					= errors.New("title is empty")
	ErrDescriptionInvalid 			= errors.New("description is empty")
	ErrPriceInvalid 				= errors.New("price is empty")
	ErrAmountInvalid 				= errors.New("amount is empty")
	ErrCurrencyCodeInvalid 			= errors.New("currencyCode is empty")
	ErrOfferTypeInvalid 			= errors.New("offerType is empty")
	ErrMeasureUnitCodeInvalid 		= errors.New("measureUnitCode is empty")
	ErrDateExpiresInvalid 			= errors.New("dateExpires is empty")
	ErrCountryInvalid 				= errors.New("country is empty")
	ErrCityInvalid 					= errors.New("city is empty")
	ErrIDInvalid 					= errors.New("id is empty")
	ErrCreatedOnInvalid 			= errors.New("created_on is empty")
	ErrOfferIDInvalid 				= errors.New("offer_id is empty")
)