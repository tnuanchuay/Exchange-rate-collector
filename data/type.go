package data

type ExchangeRates struct {
	Base	string		`json:"base"`
	Date	string		`json:"date"`
	Rate	Rates		`json:"rates"`
}

type Rates struct {
	AUD float64		`json:"AUD"`
	CAD float64		`json:"CAD"`
	CHF float64		`json:"CHF"`
	CYP float64		`json:"CYP"`
	CZK float64		`json:"CZK"`
	DKK float64		`json:"DKK"`
	EEK float64		`json:"EEK"`
	GBP float64		`json:"GBP"`
	HKD float64		`json:"HKD"`
	HUF float64		`json:"HUF"`
	ISK float64		`json:"ISK"`
	JPY float64		`json:"JPY"`
	KRW float64		`json:"KRW"`
	LTL float64		`json:"LTL"`
	LVL float64		`json:"LVL"`
	MTL float64		`json:"MTL"`
	NOK float64		`json:"NOK"`
	NZD float64		`json:"NZD"`
	PLN float64		`json:"PLN"`
	ROL float64		`json:"ROL"`
	SEK float64		`json:"SEK"`
	SGD float64		`json:"SGD"`
	SIT float64		`json:"SIT"`
	SKK float64		`json:"SKK"`
	TRL float64		`json:"TRL"`
	USD float64		`json:"USD"`
	ZAR float64		`json:"ZAR"`

}
