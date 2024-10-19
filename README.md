# Donation Form

A simple donation form using Stripe for payment processing, built with HTML, Bootstrap, and JavaScript. Users can enter their details and donate using Stripe's Checkout API.

## Features

- Donation form with fields for amount, name, address, zip code, city, email, and phone number.
- Integration with Stripe to securely process payments.
- Redirection to Stripe Checkout for handling donations.

## Prerequisites

Before you start, make sure you have the following installed:

- A Stripe account with API keys. You can use the provided test API key (`pk_test_...`), but replace it with your own from the [Stripe Dashboard](https://dashboard.stripe.com/test/apikeys).
- Golang installed
  
## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/donation-form.git
cd donation-form
````

### 2. Run the code
```bash
go run main.go
````

### 3. Try it on the web
Open the index.html file and try to make a donation. You should use one of the sample credit card numbers Stripe provides.

