import React from 'react';
import axios from 'axios';
import PortfolioItem from './PortfolioItem'

export default class Portfolio extends React.Component {
 
  constructor(props) {
    super(props);
  }

  async componentDidMount(){

    // we only want to load the portfolio and transaction
    // data once- when we log in and navigate to this window

    // after that, the client application will update its state
    // based on the transactions that clear, not by hitting the API
    // each time to get the data

    if (!this.props.hasLoadedData) {

      let portfolio
      let transactionHistory

      try {

        let theHeader = { 
                          headers: {
                            "Authorization": 'Bearer ' + this.props.profile.token 
                          }
                        }
        Promise.all( // allows you to run api calls in parallel rather than serially
          [ 
            ( portfolio = await axios.post('/getportfolio', {}, theHeader) ),
            ( transactionHistory = await axios.post('/getallTransactions', {}, theHeader) )
          ]
        )
      }
      catch(error){

        console.log(error)
        // alert(error.response.data.message)
      }

      this.props.loadInitialData(JSON.parse(portfolio.data), JSON.parse(transactionHistory.data))
    }

  }

  render() {

    return (
      <div>

          { this.props.portfolio !== null ? 
              Object.entries(this.props.portfolio).map( (item, i) =>  <PortfolioItem key={i} data={item[1]} /> ) : ''
          }

      </div>
    );
  }
}