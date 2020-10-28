package model

type Offer struct {
	OfferID   			string    	`json:"offer_id" example:"1d586b05-7b80-4a3a-bf2c-ce48169d4e85" format:"string"`
	Title   			string    	`json:"title" example:"Сгущенка" format:"string"`
	Description 		string 		`json:"description" example:"Оригинальная сгущенка Рогачев" format:"string"`
	Price				string 		`json:"price" example:"1000000" format:"int"`
	Amount 				string 		`json:"amount" example:"100" format:"int"`
	CurrencyCode 		string 		`json:"currency_code" example:"RUB" format:"string"`
	OfferType 			string 		`json:"offer_type" example:"BUY" format:"string"`
	MeasureUnitCode 	string 		`json:"measure_unit_code" example:"KG" format:"string"`
	DateExpires 		string 		`json:"date_expires" example:"2020-10-28T14:58:56.059Z" format:"string"`
	Country 			string 		`json:"country" example:"Российская Федерация" format:"string"`
	City 				string 		`json:"city" example:"Москва" format:"string"`
	CreatedOn			string 		`json:"created_on" example:"2020-10-28T14:58:56.059Z" format:"string"`
}