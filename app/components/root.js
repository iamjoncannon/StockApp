import React from 'react'
import LogIn  from './Login'
import SignUp from './SignUp'
import Portfolio from './Portfolio'
import TransactionHistory from './TransactionHistory'
import MakeTrade from './MakeTrade'
import Socket from './Socket'
import { asyncLogInCall, asyncSignUpCall, asyncMakeTrade, asyncGetOpeningPrice } from './asyncCalls'
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import TabContainer from './DashTab'

export default class Root extends React.Component {

	constructor(props) {
	    super(props);
	    this.state = {
	    	isLoggedIn: false,
	    	hasLoadedData: false,
	    	
	    	page: 'login',
	    	tab: 0,
	    	
	    	socket: null,
	    	
	    	profile: null, // Balance, Name, token
	    	portfolio: {1:{Symbol: "", Quantity: "", Price: ""}},
	    	transactionHistory: {1:{Symbol: "", Quantity: "", Date: ""}},
	    	openingPriceCache: {},
	    	cachedPriceList: {},
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

		let updatedOpeningPriceCache = {... openingPriceCache}

		for(let stock in repopulatedPortfolio){
			
			if( cachedPriceList ){

				if( cachedPriceList[stock] ){
				
					repopulatedPortfolio[stock]["price"] = cachedPriceList[stock]				
				}
			}

			if(!updatedOpeningPriceCache[stock]){
				
				updatedOpeningPriceCache[stock] = await asyncGetOpeningPrice(stock, this.state.profile.token)
			}
		} 

		this.setState({ portfolio: repopulatedPortfolio, 
						transactionHistory, 
						hasLoadedData: true, 
						openingPriceCache: updatedOpeningPriceCache
					  })
	}

	handleTrade = async (trade) => {

		trade.Quantity = Number(trade.Quantity)

		const data = await asyncMakeTrade(trade, this.state.profile.token)

		console.log(data)

		if(data.data.message){

			this.setState({ tradeError : data.data.message })
		}
		else{
			this.setState({ tradeError: '', 
							tab: 0, 
							profile:{...this.state.profile, Balance: data.data}
					   	 })
		}
	}

	handleSocketMessage = (stock) => {

		// continuously update price List

		let updatedPortfolio = {...this.state.portfolio } 

		let cachedPriceList = {...this.state.cachedPriceList }

		cachedPriceList[stock.symbol] = stock.price

		updatedPortfolio[stock.symbol]["price"] = stock.price

		this.setState({portfolio: updatedPortfolio, cachedPriceList})
	}

	render(){
		
		const { portfolio } = this.state
		
		return ( 


		    <div> 

		    	{ this.state.hasLoadedData ? <Socket portfolio={portfolio} 
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

			    	<div>
				      <AppBar position="static">
				          <Tab label={this.state.profile.Name + "   Balance: $" + this.state.profile.Balance} />
				        <Tabs value={this.state.tab} onChange={(x, y)=> this.setState({tab: y})}>
				          <Tab label="Portfolio" />
				          <Tab label="Trading History" />
				          <Tab label="Make a Trade" />
				        </Tabs>
				      </AppBar>

				      {this.state.tab === 0 && <TabContainer> 
				      						
				      								<Portfolio 
				      									loadPortfolioData={this.loadPortfolioData} 
									    		   	   	portfolio={ portfolio }
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
				      									Balance={this.state.profile.Balance} 
				      									portfolio={ portfolio }
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
