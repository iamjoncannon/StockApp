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
	    	portfolio: {1:{Symbol: "", Quantity: "", Price: ""}},
	    	transactionHistory: {1:{Symbol: "", Quantity: "", Date: ""}},
	    	socket: null,
	    	currentPrice: {},
	    	cachedPriceList: {},
	    	isDataLoaded: false
	    }
	}

	handleSignUp = async ({ firstName, lastName, email, password }) => {

		const name = firstName + lastName

		const data = await asyncSignUpCall(name, email, password)

		// console.log(data)
		alert("Please sign in with your credentials")
		location.reload()
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

		let { cachedPriceList } = this.state

		let repopulatedPortfolio = {...portfolio }

		for(let stock in repopulatedPortfolio){
			
			if( cachedPriceList ){

				if( cachedPriceList[stock] ){
					// console.log(cachedPriceList[stock])
					repopulatedPortfolio[stock]["price"] = cachedPriceList[stock]				
				}
			}
		} 

		this.setState({ portfolio: repopulatedPortfolio, transactionHistory, isDataLoaded: true })
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

		return ( 

		    <div> 

		    	{ this.state.isDataLoaded ? <Socket portfolio={this.state.portfolio} 
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
