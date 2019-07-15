import React from 'react';
import MenuItem from '@material-ui/core/MenuItem';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import Select from '@material-ui/core/Select';
import { asyncGetOnePrice } from './asyncCalls'
import allSymbols from './symbolHash.json'

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

  componentDidMount(){
  
  }

  componentDidUpdate(){


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

    let formComplete = allSymbols[this.state.Symbol] && filled(Type) && filled(this.state.Symbol) && filled(Quantity)

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

      Price * Quantity  <= this.props.Balance 
    }

    return formComplete && canCoverPurchase && canCoverSale
  }

  currentHoldingsMessage = () => {

    return "You currently hold " + this.props.portfolio[this.state.Symbol].quantity + ' shares of this stock.'
  }

  render() {

    return (

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

      { this.validateTrade() ? 

        <Button
          onClick={this.defaultPreventer}
          fullWidth
          variant="contained"
          color="primary"
          className={"blank"}
        >

          {"Make Trade"}

        </Button>
      : ""}
      <div>
        { allSymbols[this.state.Symbol] }
      </div>
      <div>
        { this.props.portfolio[this.state.Symbol] ? this.currentHoldingsMessage() : 
          'You currently hold 0 shares of this stock.'}
      </div>
      <div>

        { this.props.tradeError ? 

          `${this.props.tradeError}` : ''
        }
      </div>
    </form>  
    );
  }
}

function filled(field){
      return !(field === '')
}
