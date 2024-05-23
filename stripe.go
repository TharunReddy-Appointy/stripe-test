//package main
//
//import (
//	"fmt"
//	"github.com/stripe/stripe-go/v72"
//	"github.com/stripe/stripe-go/v72/paymentintent"
//	"github.com/stripe/stripe-go/v72/setupintent"
//	"github.com/stripe/stripe-go/v72/sub"
//)
//
//func main() {
//	// Set your secret key
//	stripe.Key = "sk_test_51L2pmSSAAtniLL2Y6tce9PlfWTM0fm3wsFLX2bKu4JC3kQiwFg1JtLDKnAYugCsSKobDmNFzOxrodRcyPagz4TGG007TDO6xlj"
//
//	//Create Setup Intent with Mandate
//	setupIntent, err := createSetupIntentWithMandate()
//	if err != nil {
//		fmt.Printf("Error creating Setup Intent: %v\n", err)
//		return
//	}
//	fmt.Printf("Setup Intent created: %v\n", setupIntent)
//
//	// Create Payment Intent with Mandate
//	paymentIntent, err := createPaymentIntentWithMandate()
//	if err != nil {
//		fmt.Printf("Error creating Payment Intent: %v\n", err)
//		return
//	}
//	fmt.Printf("Payment Intent created: %v\n", paymentIntent)
//
//	//Create Subscription with Mandate
//	subscription, err := createSubscriptionWithMandate()
//	if err != nil {
//		fmt.Printf("Error creating Subscription: %v\n", err)
//		return
//	}
//	fmt.Printf("Subscription created: %v\n", subscription)
//}
//
//func createSetupIntentWithMandate() (*stripe.SetupIntent, error) {
//	params := &stripe.SetupIntentParams{
//		Customer:           stripe.String("cus_Q4YvImMRbyMQJ6"),
//		Usage:              stripe.String(string(stripe.SetupIntentUsageOffSession)),
//		PaymentMethodTypes: []*string{stripe.String("card")},
//		PaymentMethod:      stripe.String("pm_1PEPnISAAtniLL2YuTK2SNvy"),
//		PaymentMethodOptions: &stripe.SetupIntentPaymentMethodOptionsParams{
//			Card: &stripe.SetupIntentPaymentMethodOptionsCardParams{
//				MandateOptions: &stripe.SetupIntentPaymentMethodOptionsCardMandateOptionsParams{
//					Reference:      stripe.String("Dry Fruits Recurring"),
//					Description:    stripe.String("Dry Fruits Recurring"),
//					Amount:         stripe.Int64(2000000),
//					Currency:       stripe.String(string(stripe.CurrencyINR)),
//					StartDate:      stripe.Int64(1675238400),
//					AmountType:     stripe.String(string(stripe.SetupIntentPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum)),
//					Interval:       stripe.String(string(stripe.SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalMonth)),
//					IntervalCount:  stripe.Int64(1),
//					SupportedTypes: []*string{stripe.String("india")},
//				},
//			},
//		},
//	}
//	intent, err := setupintent.New(params)
//	if err != nil {
//		return nil, err
//	}
//	return intent, nil
//}
//
//func createPaymentIntentWithMandate() (*stripe.PaymentIntent, error) {
//	//params := // Set your secret key. Remember to switch to your live secret key in production.
//	// See your keys here: https://dashboard.stripe.com/apikeys
//	stripe.Key = "sk_test_51L2pmSSAAtniLL2Y6tce9PlfWTM0fm3wsFLX2bKu4JC3kQiwFg1JtLDKnAYugCsSKobDmNFzOxrodRcyPagz4TGG007TDO6xlj"
//
//	params := &stripe.PaymentIntentParams{
//		Amount:             stripe.Int64(2000000),
//		Currency:           stripe.String(string(stripe.CurrencyINR)),
//		Customer:           stripe.String("cus_Q4YvImMRbyMQJ6"),
//		SetupFutureUsage:   stripe.String(string(stripe.PaymentIntentSetupFutureUsageOffSession)),
//		PaymentMethodTypes: []*string{stripe.String("card")},
//		PaymentMethod:      stripe.String("pm_1PEPnISAAtniLL2YuTK2SNvy"),
//		PaymentMethodOptions: &stripe.PaymentIntentPaymentMethodOptionsParams{
//			Card: &stripe.PaymentIntentPaymentMethodOptionsCardParams{
//				MandateOptions: &stripe.PaymentIntentPaymentMethodOptionsCardMandateOptionsParams{
//					Reference:      stripe.String("Dry Fruits Recurring"),
//					Description:    stripe.String("Dry Fruits Recurring"),
//					Amount:         stripe.Int64(2000000),
//					StartDate:      stripe.Int64(1675238400),
//					AmountType:     stripe.String(string(stripe.PaymentIntentPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum)),
//					Interval:       stripe.String(string(stripe.PaymentIntentPaymentMethodOptionsCardMandateOptionsIntervalSporadic)),
//					IntervalCount:  stripe.Int64(1),
//					SupportedTypes: []*string{stripe.String("india")},
//				},
//			},
//		},
//	}
//	intent, err := paymentintent.New(params)
//	if err != nil {
//		return nil, err
//	}
//	return intent, nil
//}
//
//func createSubscriptionWithMandate() (*stripe.Subscription, error) {
//	params := &stripe.SubscriptionParams{
//		Customer: stripe.String("cus_PzhJ6cIyUkmJIL"),
//		Items: []*stripe.SubscriptionItemsParams{
//			{
//				Price:    stripe.String("price_1P8dTcSAAtniLL2Y8NHoPc1r"),
//				Quantity: stripe.Int64(1),
//			},
//		},
//		DefaultPaymentMethod: stripe.String("pm_1P9joVSAAtniLL2YUaFibAO0"),
//		PaymentSettings: &stripe.SubscriptionPaymentSettingsParams{
//			PaymentMethodOptions: &stripe.SubscriptionPaymentSettingsPaymentMethodOptionsParams{
//				Card: &stripe.SubscriptionPaymentSettingsPaymentMethodOptionsCardParams{
//					MandateOptions: &stripe.SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsParams{
//						Amount:      stripe.Int64(10000),
//						AmountType:  stripe.String(string(stripe.PaymentIntentPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum)),
//						Description: stripe.String("Dry Fruits Recurring"),
//					},
//					RequestThreeDSecure: stripe.String("automatic"),
//				},
//			},
//		},
//	}
//
//	sub, err := sub.New(params)
//	if err != nil {
//		return nil, err
//	}
//	return sub, nil
//}

// -----------comment------------
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/setupintent"
	"html/template"
	"net/http"

	"github.com/stripe/stripe-go/v72"
)

func main() {
	// Set your secret key
	stripe.Key = "sk_test_51L2pmSSAAtniLL2Y6tce9PlfWTM0fm3wsFLX2bKu4JC3kQiwFg1JtLDKnAYugCsSKobDmNFzOxrodRcyPagz4TGG007TDO6xlj"

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/create-customer", createCustomerHandler).Methods("POST")
	r.HandleFunc("/create-setup-intent", createSetupIntentHandler).Methods("POST")
	r.HandleFunc("/create-subscription", createSubscriptionHandler).Methods("POST")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func createSetupIntentHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var reqBody struct {
		CustomerID string `json:"customer_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create setup intent with mandate options
	params := &stripe.SetupIntentParams{
		Customer: stripe.String(reqBody.CustomerID),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}
	setupIntent, err := setupintent.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the setup intent ID
	respBody := struct {
		SetupIntentID           string `json:"setup_intent_id"`
		SetupIntentClientSecret string `json:"setup_intent_client_secret"`
		CustomerId              string `json:"customer_id"`
	}{
		SetupIntentID:           setupIntent.ID,
		SetupIntentClientSecret: setupIntent.ClientSecret,
		CustomerId:              setupIntent.Customer.ID,
	}
	json.NewEncoder(w).Encode(respBody)
}

func createCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body

	//var Address struct{
	//	Postal_Code string `json:"postal_Code"`
	//
	//}
	var reqBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		//Address string `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "issue in decoding", http.StatusBadRequest)
		return
	}

	// Create customer in Stripe
	params := &stripe.CustomerParams{
		Name:  stripe.String(reqBody.Name),
		Email: stripe.String(reqBody.Email),
		Address: &stripe.AddressParams{
			PostalCode: stripe.String("123456"),
		},
	}
	c, err := customer.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the customer ID
	respBody := struct {
		CustomerID string `json:"customer_id"`
	}{
		CustomerID: c.ID,
	}
	json.NewEncoder(w).Encode(respBody)
}

func createSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var reqBody struct {
		CustomerID    string `json:"customer_id"`
		PaymentMethod string `json:"payment_method"`
		PriceID       string `json:"price_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Create setup intent with mandate options
	//paramsSetupIntent := &stripe.SetupIntentParams{
	//	Customer:           stripe.String(reqBody.CustomerID),
	//	Usage:              stripe.String(string(stripe.SetupIntentUsageOffSession)),
	//	PaymentMethodTypes: []*string{stripe.String("card")},
	//	PaymentMethod:      stripe.String(reqBody.PaymentMethod),
	//	PaymentMethodOptions: &stripe.SetupIntentPaymentMethodOptionsParams{
	//		Card: &stripe.SetupIntentPaymentMethodOptionsCardParams{
	//			MandateOptions: &stripe.SetupIntentPaymentMethodOptionsCardMandateOptionsParams{
	//				Reference:      stripe.String(reqBody.CustomerID),
	//				Description:    stripe.String("Dry Fruits Recurring"),
	//				Amount:         stripe.Int64(2000000),
	//				Currency:       stripe.String(string(stripe.CurrencyINR)),
	//				StartDate:      stripe.Int64(1675238400),
	//				AmountType:     stripe.String(string(stripe.SetupIntentPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum)),
	//				Interval:       stripe.String(string(stripe.SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalMonth)),
	//				IntervalCount:  stripe.Int64(1),
	//				SupportedTypes: []*string{stripe.String("india")},
	//			},
	//		},
	//	},
	//}
	//intent, err := setupintent.New(paramsSetupIntent)
	//if err != nil {
	//	http.Error(w, "setup intent error", http.StatusInternalServerError)
	//}
	//
	//params := &stripe.SetupIntentConfirmParams{
	//	PaymentMethod: stripe.String(intent.PaymentMethod.ID),
	//}
	//result, err := setupintent.Confirm(intent.ID, params)
	//if result.Status != "succeeded" {
	//	_, _ = setupintent.Confirm(intent.ID, params)
	//}
	//fmt.Println("result: ", result)

	//Create payment intent with mandate options
	paymentIntentParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1000000),
		Currency: stripe.String(string(stripe.CurrencyINR)),
		Customer: stripe.String(reqBody.CustomerID),
		//Customer: stripe.String("cus_Q6s6WqiSjNnB9k"),
		SetupFutureUsage:   stripe.String(string(stripe.PaymentIntentSetupFutureUsageOffSession)),
		PaymentMethodTypes: []*string{stripe.String("card")},
		PaymentMethod:      stripe.String(reqBody.PaymentMethod),
		//PaymentMethod: stripe.String("pm_1PGeMxSAAtniLL2Y859Ck4Lb"),
		//Mandate:       stripe.String("mandate_1PGeN9SAAtniLL2YevYPrj6m"),
		Confirm:       stripe.Bool(true),
		CaptureMethod: stripe.String(string(stripe.PaymentIntentCaptureMethodManual)),
		PaymentMethodOptions: &stripe.PaymentIntentPaymentMethodOptionsParams{
			Card: &stripe.PaymentIntentPaymentMethodOptionsCardParams{
				MandateOptions: &stripe.PaymentIntentPaymentMethodOptionsCardMandateOptionsParams{
					Reference:      stripe.String(reqBody.CustomerID),
					Description:    stripe.String("Dry Fruits Recurring"),
					Amount:         stripe.Int64(2000000),
					StartDate:      stripe.Int64(1675238400),
					AmountType:     stripe.String(string(stripe.PaymentIntentPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum)),
					Interval:       stripe.String(string(stripe.PaymentIntentPaymentMethodOptionsCardMandateOptionsIntervalSporadic)),
					IntervalCount:  stripe.Int64(1),
					SupportedTypes: []*string{stripe.String("india")},
				},
			},
		},
	}

	paymentIntent, err := paymentintent.New(paymentIntentParams)
	if err != nil {
		http.Error(w, "payment intent error", http.StatusInternalServerError)
		return
	}

	//params := &stripe.PaymentIntentConfirmParams{
	//	PaymentMethod: stripe.String(paymentIntent.PaymentMethod.ID),
	//	ReturnURL:     stripe.String("https://www.google.com"),
	//}
	//result, err := paymentintent.Confirm(paymentIntent.ID, params)

	//Create subscription with mandate options

	//params2 := &stripe.PaymentMethodAttachParams{
	//	Customer: stripe.String(reqBody.CustomerID),
	//}
	//fmt.Println(params2.Customer)
	//result, err := paymentmethod.Attach(reqBody.PaymentMethod, params2)
	//fmt.Println(result)
	//subscriptionParams := &stripe.SubscriptionParams{
	//	Customer: stripe.String(reqBody.CustomerID),
	//	Items: []*stripe.SubscriptionItemsParams{
	//		{
	//			Price:    stripe.String(reqBody.PriceID),
	//			Quantity: stripe.Int64(1),
	//		},
	//	},
	//	DefaultPaymentMethod: stripe.String(reqBody.PaymentMethod),
	//	PaymentSettings: &stripe.SubscriptionPaymentSettingsParams{
	//		PaymentMethodOptions: &stripe.SubscriptionPaymentSettingsPaymentMethodOptionsParams{
	//			Card: &stripe.SubscriptionPaymentSettingsPaymentMethodOptionsCardParams{
	//				MandateOptions: &stripe.SubscriptionPaymentSettingsPaymentMethodOptionsCardMandateOptionsParams{
	//					Amount:      stripe.Int64(10000),
	//					AmountType:  stripe.String(string(stripe.PaymentIntentPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum)),
	//					Description: stripe.String("Dry Fruits Recurring"),
	//				},
	//				RequestThreeDSecure: stripe.String("automatic"),
	//			},
	//		},
	//	},
	//}
	//subscriptionParams.AddExpand("latest_invoice.payment_intent")
	//subscription, err := sub.New(subscriptionParams)
	//if err != nil {
	//	http.Error(w, "create sub error", http.StatusInternalServerError)
	//	return
	//}
	//params := &stripe.PaymentIntentConfirmParams{
	//	PaymentMethod: stripe.String(subscription.DefaultPaymentMethod.ID),
	//	ReturnURL:     stripe.String("https://www.google.com"),
	//}
	//result, err := paymentintent.Confirm(subscription.LatestInvoice.PaymentIntent.ID, params)

	// Respond with the setup intent, payment intent, and subscription IDs
	respBody := struct {
		SetupIntentID   string `json:"setup_intent_id"`
		PaymentIntentID string `json:"payment_intent_id"`
		SubscriptionID  string `json:"subscription_id"`
		ClientSecret    string `json:"client_secret"`
	}{
		//SetupIntentID: intent.ID,
		PaymentIntentID: paymentIntent.ClientSecret,
		//SubscriptionID: subscription.ID,
		ClientSecret: paymentIntent.ClientSecret,
	}
	json.NewEncoder(w).Encode(respBody)
}
