items to fix:

1. implement the color changing thing

2. trade validation- pass portfolio into make trade and have it display the current holdings when a stock is selected, if buy, validate with balance, if sale, validate with holding

3. populate balance from trade completion

4. implement redis cache for external api calls

think about using this for trade screen:

https://material-ui.com/getting-started/page-layout-examples/checkout/



USER STORIES:

1. sign up with name, email password
- 5k in account
- user only registers once

2. login with email and password

3. buy stocks at current price with symbol and number of shares
- validates balance
- validates stock symbol

4. lists transactions 

5. lists all stocks owned with current values
- values = latest price and quantity owned

6.  color of stock in portfolio changes during the day
	- red if lost against opening price
	- grey if same
	- green if gained against opening price


