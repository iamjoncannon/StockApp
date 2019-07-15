
// define interface with API 
// as separate concern from UI components

import axios from 'axios';

const makeHeader = (token) => {
  
  return { headers: {"Authorization": 'Bearer ' + token}}
}

export const asyncLogInCall = async (email, password) => {

  let res 
  try{
    res = await axios.post('/login', { email, password} )    
  }
  catch(error){
    let  duplicateEmail = error.response.data.message === `pq: duplicate key value violates unique constraint "unique_email"`
    if (duplicateEmail){
      alert("This Email is already on file, please try again with a unique email.")
    }
    else{
      alert(error.response.data.message)
    }
  }
  
  let parsed = JSON.parse(res.data)

  console.log(parsed)
  
  return { Name: parsed.Name, email, token: parsed.token }
}

export const asyncSignUpCall = async (Name, email, password) => {
    
  const signUpInfo = { Name, email, password }
  let token
  try {
    token = await axios.post('/signup', signUpInfo )
  }
  catch(error){
    let  duplicateEmail = error.response.data.message === `pq: duplicate key value violates unique constraint "unique_email"`
    if (duplicateEmail){
      alert("This Email is already on file, please try again with a unique email.")
    }
    else{
      console.log(error.response.data.message)
      alert(error.response.data.message)
    }
  }
  signUpInfo.password = null 
  let returnedToken = token.data.token
  return { signUpInfo, returnedToken }
}

export const asyncPopulateData = async (token, callback) => {

  let portfolio
  let transactionHistory

  try {

    Promise.all( // allows us to run api calls in parallel rather than serially
      [ 
        ( portfolio = await axios.post('/getportfolio', {}, makeHeader(token)) ),
        ( transactionHistory = await axios.post('/getallTransactions', {}, makeHeader(token)) )
      ]
    )
  }
  catch(error){
    console.log(error)
    // alert(error.response.data.message)
  }

  if(portfolio.data && transactionHistory.data){
  
    callback(JSON.parse(portfolio.data), JSON.parse(transactionHistory.data))
  }
  else{
  
    callback({}, {1:{Symbol: "", Quantity: "", Date: ""}})
  }
}

export const asyncMakeTrade = async (trade, token) => {

  let data

  try {

    data = await axios.post('/maketransaction', trade, makeHeader(token)) 
  }
  catch(error){
    console.log(error)
    // alert(error.response.data.message)
  }

  return data
}

export const asyncGetOnePrice = async (symbol) => {

  const url = `https://api.iextrading.com/1.0/tops/last?symbols=${symbol}`
  
  let { data } = await axios.get(url)

  return data[0].price
}

export const asyncGetOpeningPrice = async (symbol) => {

  const url = `https://api.iextrading.com/1.0/deep/official-price?symbols=${symbol}`
  
  let { data } = await axios.get(url)

  return data[0].price
}

