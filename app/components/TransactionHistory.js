import React from 'react';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

export default class TransactionHistory extends React.Component {
 
  constructor(props) {
    super(props);
    this.state = {

    }
  }

  render() {

    console.log(this.props.transactionHistory)

    return (

      <div>

          { this.props.portfolio !== null ? 

          <Paper className={"blank"}>
            <Table className={"blank"}>
              <TableHead>
                <TableRow>
                  <TableCell>Symbol</TableCell>
                  <TableCell align="right">Date Trade Conducted</TableCell>                  
                  <TableCell align="right">Transaction Type</TableCell>                  
                  <TableCell align="right">Price at Previous Trade</TableCell>                  
                  <TableCell align="right">Shares Traded</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {Object.entries(this.props.transactionHistory).reverse().map((row, i) => (
                  <TableRow key={i}>
                    <TableCell component="th" scope="row">
                      {row[1].Symbol}
                    </TableCell>
                    <TableCell align="right">{row[1].Date}</TableCell>
                    <TableCell align="right">{row[1].Type}</TableCell>
                    <TableCell align="right">{row[1].Price}</TableCell>
                    <TableCell align="right">{row[1].Quantity}</TableCell>
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
