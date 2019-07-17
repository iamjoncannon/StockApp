import React from 'react';
import { asyncPopulateData } from './asyncCalls'
import ColoredStock from './ColoredStock'
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

    const { token } = this.props.profile
    asyncPopulateData(token, this.props.loadPortfolioData)
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
                  <TableCell align="right">Opening Price USD</TableCell>                  
                  <TableCell align="right">Current Holdings USD</TableCell>
                  <TableCell align="right">Current Price USD</TableCell>                  
                  <TableCell align="right">Current Total Value USD </TableCell>                  
                </TableRow>
              </TableHead>
              <TableBody>
                {Object.entries(this.props.portfolio).filter(row => row[1].quantity > 0).map((row, i) => (
                  <TableRow key={i}>
                    <TableCell component="th" scope="row">
                      <ColoredStock cell={row[1]} openingPrice={this.props.openingPriceCache[row[1].symbol]}/>
                    </TableCell>
                    <TableCell align="right">{this.props.openingPriceCache[row[1].symbol]}</TableCell>
                    <TableCell align="right">{row[1].quantity}</TableCell>
                    <TableCell align="right">{row[1].price}</TableCell>
                    <TableCell align="right">{row[1].price * row[1].quantity}</TableCell>
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