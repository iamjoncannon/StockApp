import React from 'react';
import symbolHash from './symbolHash.json'

export default class Socket extends React.Component {
 
  constructor(props) {
    super(props);
 
  }

  connectToSocket = () => {

    const socket = io('https://ws-api.iextrading.com/1.0/last')
    const myBook = []
    
    for (let stock in this.props.portfolio){

      myBook.push(this.props.portfolio[stock].symbol)
    }

    socket.on('connect', () => {

      myBook.forEach( stock=>{
        socket.emit('subscribe', stock)
      })
    })

    socket.on('message', message => {
    
      this.props.handleSocketMessage(JSON.parse(message))
    })

    this.setState({ socket, portfolioSize: Object.keys(this.props.portfolio).length })
  }


  async componentDidUpdate(){

    let currentPortfolio = Object.keys(this.props.portfolio).length

    let portfolioChanged = currentPortfolio !== this.state.portfolioSize

    if( portfolioChanged ){
      // otherwise we get a memory leak
      await this.state.socket.close()
      this.connectToSocket()
    }
  }

  componentDidMount(){
    this.connectToSocket()
  }

  render() {

    return (
      <div></div>
    );
  }
}