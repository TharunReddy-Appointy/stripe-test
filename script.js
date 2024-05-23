var stripe = Stripe('pk_test_51L2pmSSAAtniLL2YKyvdcY7ofjGiuMFcXP8389AfB57aaUORpYEyaE8coMi8QLFAfcHcvHI5xPO41e2XN9aeXsy000q90zxlSX');

// Create an instance of Elements
var elements = stripe.elements();

// Custom styling can be passed to options when creating an Element
var style = {
    base: {
        // Add your base input styles here. For example:
        fontSize: '16px',
        color: '#32325d',
    },
};

// Create an instance of the card Element
var card = elements.create('card', { style: style });

// Add an instance of the card Element into the `card-element` div
card.mount('#card-element');

// Handle form submission
var form = document.getElementById('payment-form');
form.addEventListener('submit', function(event) {
    event.preventDefault();

    // Disable the submit button to prevent multiple submissions
    document.getElementById('submit').disabled = true;

    // Create a customer
    fetch('/create-customer', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            name: document.getElementById('name').value,
            email: document.getElementById('email').value,
        }),
    })
        .then(function(response) {
            console.log(response)
            // return response.json();
        })
        .then(function(data) {
            // Create payment method with mandate options
            stripe.createPaymentMethod({
                type: 'card',
                card: card,
                billing_details: {
                    name: document.getElementById('name').value,
                    email: document.getElementById('email').value,
                },
            }).then(function(result) {
                if (result.error) {
                    // Inform the user if there was an error
                    var errorElement = document.getElementById('card-errors');
                    errorElement.textContent = result.error.message;

                    // Enable the submit button
                    document.getElementById('submit').disabled = false;
                } else {
                    console.log('else')
                    // Otherwise, proceed with subscription creation
                    fetch('/create-subscription', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify({
                            customer_id: data.customer_id,
                            payment_method: result.paymentMethod.id,
                            price_id: 'price_1P8dTcSAAtniLL2Y8NHoPc1r', // Replace with your price ID
                        }),
                    }).then(function(response) {
                        return response.json();
                    }).then(function(data) {
                        // Handle the subscription creation response
                        console.log(data);
                    }).catch(function(error) {
                        console.error('Error:', error);

                        // Enable the submit button
                        document.getElementById('submit').disabled = false;
                    });
                }
            });
        })
        .catch(function(error) {
            console.error('Error:', error);

            // Enable the submit button
            document.getElementById('submit').disabled = false;
        });
});
