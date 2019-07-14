import React from 'react';
import axios from 'axios';
import PortfolioItem from './PortfolioItem'
import {asyncPopulateData} from './asyncCalls'

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

      const { token } = this.props.profile

      asyncPopulateData(token, this.props.loadInitialData)
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