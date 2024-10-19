package types

type Donation struct {
	Amount  int64  `json:"amount"`
	Name    string `json:"name"`
	Address string `json:"address"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
