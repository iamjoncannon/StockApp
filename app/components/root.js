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
	    	socket: null,
	    	currentPrice: {},
	    	openingPrice: {}
	    }
	}

	handleSignUp = async ({ firstName, lastName, email, password }) => {

		const name = firstName + lastName

		const data = await asyncSignUpCall(name, email, password)

		// console.log(data)

		this.setState({ profile: data.signUpInfo, 
						token: data.returnedToken, 
						isLoggedIn: true
					})
	}

	handleLogIn = async (email, password) => {

		const data = await asyncLogInCall(email, password)

		this.setState({ profile: { name: data.name, 
								   email: data.email,
								   token: data.token,
								   Balance: data.Balance
								 }, 
						isLoggedIn: true,
						page: 'portfolio'
					})
	}

	loadPortfolioData = (portfolio, transactionHistory) => {

		// when we load data here, check if its in the 
		// current price list
		// if so, append to the portfolio entry

		this.setState({portfolio, transactionHistory})
	}

	handleTrade = async (trade) => {

		trade.Quantity = Number(trade.Quantity)

		const data = await asyncMakeTrade(trade, this.state.profile.token)

		// console.log(data.data.message)

		if(data.data.message){

			this.setState({ tradeError : data.data.message })
		}
		else{
			this.setState({tradeError: '', tab: 0})
		}
	}

	handleSocketMessage = (stock) => {

		// here we will continuously 

		// let openingPriceList = this.state.portfolio
		let openingPriceList = {}

		console.log(stock.symbol, stock.price)

		openingPriceList[stock.symbol]["price"] = stock.price
		openingPriceList[stock.symbol] = stock.price

		this.setState({openingPriceList})
	}

	render(){

		// console.log(this.state)

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
				      
				      {this.state.tab === 2 && <TabContainer> 
				      								<MakeTrade 
				      									handleTrade={this.handleTrade}
				      									tradeError={this.state.tradeError}
				      	   						    /> 
				      							</TabContainer>}
				    </div>
		    	}
		    </div>
		 )
	}
}
