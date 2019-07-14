import React from 'react';

export default class MakeTrade extends React.Component {
 
  constructor(props) {
    super(props);
    this.state = {

    }
  }

  componentDidMount(){

  }

  handleSubmit = async (event) => {

    event.preventDefault()

    console.log(event.target.quantity.value, event.target.stock.value)


  }

  render() {

    console.log('hitting make trade')

    return (

        <div className={'container'}>

        <div>

          <form onSubmit={this.handleSubmit} autoComplete="off">

          <div>

            <label htmlFor="quantity">

              <small>Quantity</small>
            
            </label>
          
            <input
              name="quantity"
              type="quantity"
              type="number"
            />
          </div>
          
          <div>
          
            <label htmlFor="stock">
              <small>Stock Symbol</small>
            </label>
          
            <input
              name="stock"
              type="stock"
              required
            />
          </div>
          
          <div>
            <button type="submit">SUBMIT</button>
          </div>    
        </form>
        </div>
        
      </div>
    );
  }
}