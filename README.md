# TTP-FP

This repo is my stock portfolio app for TTP. 

The front end web application uses React as well as Socket.io to establish a web socket connection to the IEX data service.

The backend server is written in Golang, uses JWT (JSON Web Tokens) to implement user authentication. The user's id is placed inside the JWT and decrypted to complete processing API requests to the server. The database used is PostgreSQL, and the trading transaction is completed using prepared statements to ensure atomicity across multiple database calls.  

In order to complete the last user story, data for the opening price must be obtained from the IEX Cloud, an API that requires authentication. In order to prevent the client from calling the service with my private token, I established an API endpoint in my server, which then called the IEX endpoint with my token, and relayed the data back to the client.  In a production environment, this data from the exchange would be cached to prevent duplicate API calls, as the opening price does not fluctuate during the day.  

## USER STORIES:

1. Sign up with name, email password

2. Login with email and password 

3. Buy stocks at current price with symbol and number of shares 

4. View list of transactions 

5. View portfolio of all stocks owned with current values

6. On Portfolio, color of stocks changes during the day based on current value against opening value
