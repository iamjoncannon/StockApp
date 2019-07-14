import React from 'react';
import axios from 'axios';
import { asyncPopulateData } from './asyncCalls'

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

export default class Portfolio extends React.Component {
 
  constructor(props) {
    super(props);

    this.state = {
      rows: null
    }
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

          <Paper className={"blank"}>
            <Table className={"blank"}>
              <TableHead>
                <TableRow>
                  <TableCell>Symbol</TableCell>
                  <TableCell align="right">Current Holdings</TableCell>
                  <TableCell align="right">Current Value</TableCell>                  
                </TableRow>
              </TableHead>
              <TableBody>
                {Object.entries(this.props.portfolio).map(row => (
                  <TableRow key={row.name}>
                    <TableCell component="th" scope="row">
                      {row[1].symbol}
                    </TableCell>
                    <TableCell align="right">{row[1].quantity}</TableCell>
                    <TableCell align="right">{row[1].price}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </Paper>
          : ''
        }
      </div>
    );
  }
}