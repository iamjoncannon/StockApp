import React from 'react'
import LogIn  from './loginMUI'
import SignUp from './SignUpMUI'
import Portfolio from './Portfolio'
import TransactionHistory from './TransactionHistory'
import MakeTrade from './MakeTrade'
import Socket from './Socket'
import { asyncLogInCall, asyncSignUpCall, asyncMakeTrade, asyncGetOpeningPrice } from './asyncCalls'
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';
import TabContainer from './DashTab'

// this is the root of the application, that defines the management
// of browser application state based on user input

// as the "controller" layer, the appearance of the elements is 
// separated into other components, which are held in other files

export default class Root extends React.Component {

	constructor(props) {
	    super(props);
	    this.state = {
	    	profile: null,
	    	isLoggedIn: false,
	    	page: 'login',
	    	tab: 0,
	    	portfolio: {1:{Symbol: "", Quantity: "", Price: ""}},
	    	transactionHistory: {1:{Symbol: "", Quantity: "", Date: ""}},
	    	socket: null,
	    	currentPrices: {},
	    	openingPriceCache: {},
	    	cachedPriceList: {},
	    	hasLoadedData: false
	    }
	}

	handleSignUp = async ({ firstName, lastName, email, password }) => {

		const name = firstName + lastName

		const data = await asyncSignUpCall(name, email, password)

		alert("Please sign in with your credentials")
		location.reload()
	}

	handleLogIn = async (email, password) => {

		const data = await asyncLogInCall(email, password)

		this.setState({ profile: JSON.parse(data), 
						isLoggedIn: true,
						page: 'portfolio'
					})
	}

	loadPortfolioData = async (portfolio, transactionHistory) => {

		// when we load data here, check if its in the 
		// current price list
		// if so, append to the portfolio entry

		let { cachedPriceList, openingPriceCache } = this.state

		let repopulatedPortfolio = {...portfolio }

		let updatedOpeningPriceCach = {... openingPriceCache}

		for(let stock in repopulatedPortfolio){
			
			if( cachedPriceList ){

				if( cachedPriceList[stock] ){
					// console.log(cachedPriceList[stock])
					repopulatedPortfolio[stock]["price"] = cachedPriceList[stock]				
				}
			}

			if(!updatedOpeningPriceCach[stock]){
				updatedOpeningPriceCach[stock] = await asyncGetOpeningPrice(stock, this.state.profile.token)
			}

		} 

		this.setState({ portfolio: repopulatedPortfolio, 
						transactionHistory, 
						hasLoadedData: true, 
						openingPriceCache: updatedOpeningPriceCach
					  })
	
	}

	handleTrade = async (trade) => {

		trade.Quantity = Number(trade.Quantity)

		const data = await asyncMakeTrade(trade, this.state.profile.token)

		if(data.data.message){

			this.setState({ tradeError : data.data.message })
		}
		else{
			this.setState({tradeError: '', tab: 0})
		}
	}

	handleSocketMessage = (stock) => {

		// continuously update previous price List

		let updatedPortfolio = {...this.state.portfolio } 

		let cachedPriceList = {...this.state.cachedPriceList }

		cachedPriceList[stock.symbol] = stock.price

		updatedPortfolio[stock.symbol]["price"] = stock.price

		this.setState({portfolio: updatedPortfolio, cachedPriceList})
	}

	render(){

		// console.log(this.state)

		return ( 

		    <div> 

		    	{ this.state.hasLoadedData ? <Socket portfolio={this.state.portfolio} 
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
					<div>    

				

			    	<div className={"blank"}>
				      <AppBar position="static">
				          <Tab label={this.state.profile.Name + " Balance: " + this.state.profile.Balance} />
				        <Tabs value={this.state.tab} onChange={(x, y)=> this.setState({tab: y})}>
				          <Tab label="Portfolio" />
				          <Tab label="Trading History" />
				          <Tab label="Make a Trade" />
				        </Tabs>
				      </AppBar>

				      {this.state.tab === 0 && <TabContainer> 
				      						
				      								<Portfolio 
				      									loadPortfolioData={this.loadPortfolioData} 
									    		   	   	portfolio={this.state.portfolio}
									    				hasLoadedData={this.state.hasLoadedData}
									    				profile={this.state.profile}
									    				openingPriceCache={this.state.openingPriceCache}
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
		    		</div>
		    	}
		    </div>
		 )
	}
}
