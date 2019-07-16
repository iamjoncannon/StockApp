# TTP-FP

This repo is my stock portfolio app for TTP. 

The front end web application uses React as well as Socket.io to establish a web socket connection to the IEX data service.

The backend server is written in Golang, uses JWT (JSON Web Tokens) to implement user authentication, as well as PostgreSQL.  In order to complete the last user story, data for the opening price must be obtained from the IEX Cloud, an API that requires authentication. In order to prevent the client from calling the service with my private token, I established an API endpoint in my server, which then called the IEX endpoint with my token, and relayed the data back to the client.  In a production environment, this data would be cached to prevent duplicate API calls, as the opening price does not fluctuate during the day.  


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


items to fix:


4. implement redis cache for external api calls

think about using this for trade screen:

https://material-ui.com/getting-started/page-layout-examples/checkout/
