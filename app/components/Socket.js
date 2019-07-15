import React from 'react';

export default class Socket extends React.Component {
 
  constructor(props) {
    super(props);
 
  }

  connectToSocket = () => {

    let dayTimeTrading = 'https://ws-api.iextrading.com/1.0/tops'
    let afterHours = 'https://ws-api.iextrading.com/1.0/last'
    let time = (new Date()).getUTCHours()
    // let url = time > 12 & time < 21 ? dayTimeTrading : afterHours ;
    let url = afterHours
    const socket = io(url)
    const thisBook = []
    let currentPortfolio = Object.keys(this.props.portfolio).length

    // console.log("socket connecting to: ", url)
    
    for (let stock in this.props.portfolio){

      thisBook.push(this.props.portfolio[stock].symbol)
    }

    socket.on('connect', () => {

      thisBook.forEach( stock=>{

        // console.log(stock)

        socket.emit('subscribe', stock)
      })
      // console.log('subscribed to ', socket)
    })

    socket.on('message', message => {
    
      // console.log(message)
      // console.log(JSON.parse(message))
      this.props.handleSocketMessage(JSON.parse(message))
    })

    this.setState({ socket, portfolioSize: Object.keys(this.props.portfolio).length })
  }


  async componentDidUpdate(){

    let currentPortfolio = Object.keys(this.props.portfolio).length

    let portfolioChanged = currentPortfolio !== this.state.portfolioSize

    // console.log(this.state, this.props)

    if( portfolioChanged ){
      await this.state.socket.close()
      this.connectToSocket
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