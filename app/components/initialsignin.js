import React  from 'react';
import axios from 'axios';

export default class LogIn extends React.Component {
 
  constructor(props) {
    super(props);
    this.state = { 

    }
  }

  handleLogIn = async (event) => {

  	event.preventDefault()
    
    let { email, password } = event.target
  	
    email = email.value
    password = password.value
    
    let res 

    try{
      
    	res = await axios.post('/login', {email, password} )
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

    // console.log(res)

    let parsed = JSON.parse(res.data)
    
    // console.log(parsed)

    this.props.handleLogin({Name: parsed.Name, email, token: parsed.token})
  
  }

  componentDidMount(){

  }

  render() {

    return (
      <div className={'container'}>

        <div>

        	<form onSubmit={this.handleLogIn} autoComplete="off">

          <div>

            <label htmlFor="email">

              <small>Email:</small>
            
            </label>
          
            <input
              name="email"
              type="email"
              value="cuty@example.com"
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
              value="1Yadayadayada"
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

      </div>
    );
  }
}

