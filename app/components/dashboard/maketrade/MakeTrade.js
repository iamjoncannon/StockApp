import React from 'react';
import MenuItem from '@material-ui/core/MenuItem';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import Select from '@material-ui/core/Select';
import { asyncGetOnePrice } from '../../asyncCalls'
import allSymbols from '../../symbolHash.json'
import { filled } from '../../util'
import Grid from '@material-ui/core/Grid';

const initialState = {
      Symbol: '',
      Quantity: '',
      Type: 'Buy',
      ErrorStatus: ''
}

export default class MakeTrade extends React.Component {
 
  constructor(props) {
    super(props);
    this.state = initialState
  }

  defaultPreventer = (evt) => {

    evt.preventDefault()
    evt.stopPropagation()

    this.props.handleTrade(this.state)
    this.setState(initialState)
  }

  handleSymbol = async (symbol) =>{

    this.setState({Symbol: symbol})

    if(symbol.length > 0){

      let price = await asyncGetOnePrice(symbol)

      this.setState({Price: price})
    }
    else{
      this.setState({Price: 'price'})
    }
  }

  validateTrade = () => {

    if(!this.state.Symbol) return false

    const { Price, Quantity, Type } = this.state

    let formComplete = allSymbols[this.state.Symbol] && filled(Type, Quantity, this.state.Symbol)

    let canCoverSale = true

    if(Type == 'Sell'){

      if(!this.props.portfolio[this.state.Symbol]){
        return false
      }
      else{

        canCoverSale = this.props.portfolio[this.state.Symbol].quantity >= Quantity
      }
    }

    let canCoverPurchase = true

    if(Type == "Buy"){

      canCoverPurchase = Price * Quantity  <= this.props.Balance 
    }

    return formComplete && canCoverPurchase && canCoverSale
  }

  currentHoldingsMessage = () => {

    return "You currently hold " + this.props.portfolio[this.state.Symbol].quantity + ' shares of this stock.'
  }

  currentOrderPrice = () => {

    let direction = this.state.Type === "Buy" ? 'cost' : 'yeild';

    return `This trade will ${direction} $${(this.state.Quantity * this.state.Price).toFixed(2)}`
  }

  render() {

    return (

      <div>

      <form className={'blank'} noValidate autoComplete="off">
        
        <TextField
          required
          id="standard-required"
          label="Stock Symbol"
          value={this.state.Symbol}
          onChange={(e)=>{this.handleSymbol(e.target.value)}}        
          className={'blank'}
          margin="normal"
        />

        <TextField
          required
          id="standard-required"
          label="Shares to Trade"
          value={this.state.Quantity}
          onChange={(e)=>{this.setState({Quantity: e.target.value})}}        
          className={'blank'}
          margin="normal"
        />
        
        <TextField
          required
          id="filled-disabled"
          label="Current Value"
          value={this.state.Price}
          className={'blank'}
          margin="normal"
        />
        
        <Grid container>

          <FormControl className={"blank"} style={{margin: "3"}}>

            <InputLabel htmlFor="age-simple">Buy/Sell</InputLabel>
          
            <Select
              value={this.state.Type}
              onChange={ (evt)=> this.setState({Type: evt.target.value}) }
            >
          
              <MenuItem value={'Buy'}>Buy</MenuItem>
              <MenuItem value={'Sell'}>Sell</MenuItem>

            </Select>

          </FormControl>
          
        </Grid>

      </form>  


        <table className="make-trade">

            <tr>

              <th>Stock</th>

              <td>

                

                  {
                    allSymbols[this.state.Symbol] ?

                    <span> 
                      { allSymbols[this.state.Symbol] }
                    </span>
                    :
                    ""
                  }

              
              </td>

            </tr>

            <tr>

            <th>Holding</th>
                      
            <td>
            
              <div>


                { this.props.portfolio[this.state.Symbol] ? 

                  <span> 
                    {this.currentHoldingsMessage()}
                  </span>  : 
                  
                  this.state.Symbol ? 'Stock You currently hold 0 shares of this stock.' : 

                  ""

                  }

              </div>


            </td>

            </tr>

            <tr>

            
              <th>
                Size
              </th>

              <td>


                  <div>
                    
                    {this.state.Price && this.state.Quantity ? 
                  
                      <span> 
                        {this.currentOrderPrice()} 
                      </span> 

                      : 
                      ""
                      }
                
                  </div>

              </td>
            </tr>
         

            </table>

         

              <div style={{marginTop: "2vh"}}>

                <Button
                  onClick={ this.validateTrade() ? this.defaultPreventer : ""}
                  variant="contained"
                  color="primary"
                  className={"blank"}
                  disableRipple={true}
                  style={ !this.validateTrade() ? {opacity: ".5"} : {}}
                >

                  {"Make Trade"}

                </Button>
              </div>

        <div>

          { this.props.tradeError ? 

            <h2> {this.props.tradeError}</h2> : ''
          }
        </div>


      </div>

    );
  }
}
