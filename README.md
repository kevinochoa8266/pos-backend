## pos-backend

## About 

This repository is home of the Pos-Backend for stores.

## Development

One of the first things one should do is [create a Stripe Account](https://dashboard.stripe.com/register?utm_campaign=US_en_Search_Brand_StripeAccountSignup_EXA-20466828640&utm_medium=cpc&utm_source=google&ad_content=670668113714&utm_term=create%20stripe%20account&utm_matchtype=e&utm_adposition=&utm_device=c&gclid=CjwKCAiA1MCrBhAoEiwAC2d64S4MsuPgI3CgeSwH5SP_UWyRzJlH6VpRZW7RSiKEy7h9H4i9CEo-nxoC-H0QAvD_BwE)
so that you can get access to a STRIPE_API_KEY, to make the application function and work.

Next, you want to make a copy of the .env.example and add your own .env with the values you have.


Build application with Docker:

```shell
make run_container
```

Running it locally:
Ensure that you have go version 1.21.2 or higher, and then run:

```shell
go run .
```

List all available targets from the Makefile:

```
make help
```
