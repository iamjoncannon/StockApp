import React from 'react';
import MenuItem from '@material-ui/core/MenuItem';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import Select from '@material-ui/core/Select';

export default class MakeTrade extends React.Component {
 
  constructor(props) {
    super(props);
    this.state = {
      Symbol: '',
      Quantity: '',
      Type: 'buy'
    }
  }

  componentDidMount(){

  }

  defaultPreventer = (evt) => {

    evt.preventDefault()
    evt.stopPropagation()
    this.props.handleTrade(this.state)
  }

  render() {

    return (
    <form className={'blank'} noValidate autoComplete="off">
      
      <TextField
        required
        id="standard-required"
        label="Stock Symbol"
        value={this.state.Symbol}
        onChange={(e)=>{this.setState({Symbol: e.target.value})}}        
        className={'blank'}
        margin="normal"
      />

      <TextField
        required
        id="standard-required"
        label="Shares to Purchase"
        value={this.state.Quantity}
        onChange={(e)=>{this.setState({Quantity: e.target.value})}}        
        className={'blank'}
        margin="normal"
      />

      <FormControl className={"blank"}>

        <InputLabel htmlFor="age-simple">BUY/SELL</InputLabel>
        <Select
          value={this.state.Type}
          onChange={ (evt)=> this.setState({Type: evt.target.value}) }
        >
          <MenuItem value={'Buy'}>Buy</MenuItem>
          <MenuItem value={'Sell'}>Sell</MenuItem>
        </Select>
      </FormControl>

      <Button
        onClick={this.defaultPreventer}
        fullWidth
        variant="contained"
        color="primary"
        className={"blank"}
      >
        {"Make Trade"}
      </Button>

    </form>  
    );
  }
}