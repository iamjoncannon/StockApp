import React from 'react';
import TransHistoryItem from './TransHistoryItem'

export default class TransactionHistory extends React.Component {
 
  constructor(props) {
    super(props);
    this.state = {

    }
  }

  render() {

    return (
      <div>
        { this.props.transactionHistory !== null ? 
              Object.entries(this.props.transactionHistory).map( (item, i) =>  <TransHistoryItem key={i} data={item[1]} /> ) : ''
          }
      </div>
    );
  }
}