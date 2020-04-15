package models

type Product struct {
	Id				int			`json:"id,omitempty"`
	Name			string		`json:"name,omitempty"`
	Info			string		`json:"info,omitempty"`
	Price			float64		`json:"price,omitempty"`
}

