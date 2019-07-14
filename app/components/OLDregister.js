import React from 'react';
import axios from 'axios';

export default class SignUp extends React.Component {

  constructor(props) {
      super(props);
  }
 
  handleSubmit = async (event) => {

    event.preventDefault()
    
    let { firstName, lastName, email, password } = event.target
    
    email = email.value
    password = password.value
    firstName = firstName.value
    lastName = lastName.value

    const signUpInfo = { Name: firstName + " " + lastName, email, password }

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

    this.props.handleSignUp(signUpInfo, token.data.token)
  }

  render() {

    return (
      <div>

        <form onSubmit={this.handleSubmit} autoComplete="off">
          
          <div>

            <div>
            
              <label htmlFor="firstName">
                <small>First Name: </small>
              </label>
              <input name="firstName" type="text" required />
            </div>
            
            <div>
              <label htmlFor="lastName">
                <small> Last Name: </small>
              </label>
              <input name="lastName" type="text" required />
            </div>
          
          </div>

        <div>

          <label htmlFor="email">

            <small>Email:</small>
          
          </label>
        
          <input
            name="email"
            type="email"
            required
            pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$"
          />
        </div>
        
        <div>
        
          <label htmlFor="password">
            <small>Password:</small>
          </label>
        
          <input
            name="password"
            type="password"
            required
            pattern="(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}"
            title="Must contain at least one number and one uppercase and lowercase letter, and at least 8 or more characters"
          />
        </div>

        <div>
          <button type="submit">SUBMIT</button>
        </div>    
      </form>
      </div>
    );
  }
}