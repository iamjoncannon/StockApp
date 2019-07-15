import React from 'react'
import LogIn  from './loginMUI'
import SignUp from './SignUpMUI'
import Portfolio from './Portfolio'
import TransactionHistory from './TransactionHistory'
import MakeTrade from './MakeTrade'
import Socket from './Socket'
import { asyncLogInCall, asyncSignUpCall, asyncMakeTrade } from './asyncCalls'
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';
import TabContainer from './DashTab'

export default class Root extends React.Component {

	constructor(props) {
	    super(props);
	    this.state = {
	    	profile: null,
	    	isLoggedIn: false,
	    	hasLoadedData: false,
	    	page: 'login',
	    	tab: 0,
	    	portfolio: null,
	    	transactionHistory: {1:{Symbol: "", Quantity: "", Date: ""}},
	    	socket: null
	    }
	}

	handleSignUp = async ({ firstName, lastName, email, password }) => {

		const name = firstName + lastName

		const data = await asyncSignUpCall(name, email, password)

		console.log(data)

		this.setState({ profile: data.signUpInfo, 
						token: data.returnedToken, 
						isLoggedIn: true
					})
	}

	handleLogIn = async (email, password) => {

		const data = await asyncLogInCall(email, password)

		this.setState({ profile: { name: data.name, 
								   email: data.email,
								   token: data.token
								 }, 
						isLoggedIn: true,
						page: 'portfolio'
					})
	}

	loadInitialData = (portfolio, transactionHistory) => {

		this.setState({portfolio, transactionHistory, hasLoadedData: true})
	}

	handleTrade = async (trade) => {
		
		/* 

		API NEEDS: 
		'{ "ID": 1, "TYPE": "buy", "SYMBOL": "FB", "QUANTITY":10, "PRICE": 20}'
		
		*/

		// trade.Price = this.state.portfolio[trade.Symbol].price

		trade.Quantity = Number(trade.Quantity)
		console.log(trade)

		const data = await asyncMakeTrade(trade, this.state.profile.token)

	}

	handleSocketMessage = (stock) => {

		let newPortfolio = this.state.portfolio

		newPortfolio[stock.symbol]["price"] = stock.price

		this.setState(newPortfolio)
	}

	render(){

		// console.log(this.state.portfolio)

		return ( 

		    <div> 

		    	{ this.state.portfolio ? <Socket portfolio={this.state.portfolio} 
		    									 handleSocketMessage={this.handleSocketMessage}
		    							 /> : '' }


		    	{ !this.state.isLoggedIn ? 

		    		<div>
		    			
		    			{
		    				this.state.page === 'login' ? 

		    			<LogIn handleLogIn={this.handleLogIn} 
		    				   toSignUp={()=>this.setState({page: 'signUp'})}
		    			/> 
		    			: 

					    <SignUp handleSignUp={this.handleSignUp}
					    		toLogIn={()=>this.setState({page:'login'})}
					    /> 
					
		    			}

					</div>
					
					    : 

			    	<div className={"blank"}>
				      <AppBar position="static">
				        <Tabs value={this.state.tab} onChange={(x, y)=> this.setState({tab: y})}>
				          <Tab label="Portfolio" />
				          <Tab label="Trading History" />
				          <Tab label="Make a Trade" />
				        </Tabs>
				      </AppBar>

				      {this.state.tab === 0 && <TabContainer> 
				      						
				      								<Portfolio 
				      									loadInitialData={this.loadInitialData} 
									    		   	   	portfolio={this.state.portfolio}
									    				hasLoadedData={this.state.hasLoadedData}
									    				profile={this.state.profile}
									    			/>
				      						 
				      						   </TabContainer>}

				      {this.state.tab === 1 && <TabContainer> 
				      								<TransactionHistory 
									    				transactionHistory={this.state.transactionHistory}
									    			/>
				      						   </TabContainer>}
				      
				      {this.state.tab === 2 && <TabContainer> <MakeTrade handleTrade={this.handleTrade}/> </TabContainer>}
				    </div>
		    	}
		    </div>
		 )
	}
}
