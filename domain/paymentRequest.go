package domain

type PaymentRequest struct {
	ID          int32      `json:"id"`
	Amount      float64    `json:"amount"`
	Description string     `json:"description"`
	CreditCard  CreditCard `json:"creditCard"`
}

type CreditCard struct {
	Number string `json:"number"`
	Name   string `json:"name"`
	CVV    int32  `json:"cvv"`
}
