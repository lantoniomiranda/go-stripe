package server

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/checkout/session"
	"github.com/stripe/stripe-go/v80/customer"
	config "go/stripe/configs"
	"go/stripe/types"
	"log"
	"net/http"
)

func HandleDonation(w http.ResponseWriter, r *http.Request) {
	var donation types.Donation

	if err := json.NewDecoder(r.Body).Decode(&donation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stripe.Key = config.AppConfig.StripeSecretKey

	address := &stripe.AddressParams{
		Line1:      stripe.String(donation.Address),
		PostalCode: stripe.String(donation.ZipCode),
		City:       stripe.String(donation.City),
	}

	customerParams := &stripe.CustomerParams{
		Address: address,
		Email:   stripe.String(donation.Email),
		Name:    stripe.String(donation.Name),
		Phone:   stripe.String(donation.Phone),
	}

	newCustomer, err := customer.New(customerParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := &stripe.CheckoutSessionParams{
		Customer: &newCustomer.ID,
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("eur"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name:        stripe.String("Donation from " + donation.Name),
						Description: stripe.String("This is a donation to help Luis"),
					},
					UnitAmount: stripe.Int64(donation.Amount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:         stripe.String("https://www.stripe.com/"),
		CancelURL:          stripe.String("https://www.stripe.com/"),
	}

	response, err := session.New(params)
	if err != nil {
		if newCustomer.ID != "" {
			del, err := customer.Del(newCustomer.ID, nil)
			if err != nil {
				log.Printf("Failed to delete customer %s: %v", newCustomer.ID, del)
			}
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": response.ID})
}
