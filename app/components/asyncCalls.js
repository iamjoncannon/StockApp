
// define interface with API 
// to separate concern from UI components
// akin to the "thunk" pattern 

import axios from 'axios';

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

  return {Name: parsed.Name, email, token: parsed.token}
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

  console.log(token)

  let returnedToken = token.data.token

  return { signUpInfo, returnedToken }

}
