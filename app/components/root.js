import React from 'react'
import { render } from 'react-dom'
import LogIn  from './loginMUI'
import SignUp from './signUp'
import Portfolio from './Portfolio'
import TransactionHistory from './TransactionHistory'
import MakeTrade from './MakeTrade'
import Socket from './Socket'
import { asyncLogInCall } from './asyncCalls'

export default class Root extends React.Component {

	constructor(props) {
	    super(props);
	    this.state = {

	    	profile: null,
	    	isLoggedIn: false,
	    	hasLoadedData: false,
	    	page: 'portfolio',
	    	token: null,
	    	portfolio: null,
	    	transactionHistory: null,
	    	socket: null
	    }
	}

	handleSignUp = (profileValues, token) => {

		this.setState({ profile: profileValues, 
						token: token, 
						isLoggedIn: true
					})
	}

	handleLogIn = async (email, password) => {

		const data = await asyncLogInCall(email, password)

		console.log(data)

		this.setState({ profile: { name: data.name, 
								   email: data.email,
								   token: data.token
								 }, 
						isLoggedIn: true
					})
	}

	loadInitialData = (portfolio, transactionHistory) => {

		this.setState({portfolio, transactionHistory, hasLoadedData: true})
	}

	handleTrade = () => {

	}

	handleSocketMessage = (stock) => {

		let newPortfolio = this.state.portfolio

		newPortfolio[stock.symbol]["price"] = stock.price

		this.setState(newPortfolio)
	}

	render(){

		console.log(this.state)

		return ( 

		    <div> 
		    	{ this.state.portfolio ? <Socket portfolio={this.state.portfolio} 
		    									 handleSocketMessage={this.handleSocketMessage}
		    							 /> : '' }

		    	{ !this.state.isLoggedIn ? 

		    		<div>
		    			{/* 
		    				UI component calls the separate async logic with the state
		    				management callback as the first argument
		    			*/}
		    			<LogIn handleLogIn={this.handleLogIn}/> 

					    { /*<SignUp handleSignUp={this.handleSignUp}/> */}
					</div>
					    : 
				<div>

					<span onClick={()=> this.setState({page: 'portfolio'})}> Portfolio </span>
					<span onClick={()=> this.setState({page: 'transHistory'})}> Transaction History </span>
					<span onClick={()=> this.setState({page: 'trade'})}> Make A Trade </span>

					 {this.state.page === 'portfolio' ?
						
						<div>
							<Portfolio loadInitialData={this.loadInitialData} 
			    					   portfolio={this.state.portfolio}
			    					   hasLoadedData={this.state.hasLoadedData}
			    					   profile={this.state.profile}
			    			/>
			    		</div>
			    			:
			    		this.state.page === 'transHistory' ? 
			    		<div>
			    			<TransactionHistory 
			    					   transactionHistory={this.state.transactionHistory}
			    			/>
			    		</div>
			    			: 
			    		<div>
			    			<MakeTrade handleTrade={this.handleTrade}/>
			    		</div>

			    		}
		    	</div>
		    	}
		    </div>
		 )
	}
}

