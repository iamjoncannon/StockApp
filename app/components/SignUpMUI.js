import React from 'react';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';

export default class SignIn extends React.Component {

  constructor(props) {
    super(props);

    this.state = { 
      firstName:"",
      lastName:"",
      email: "",
      password: ""
    }
  }

  defaultPreventer = (evt) => {

    evt.preventDefault()
    evt.stopPropagation()
    this.props.handleSignUp(this.state)
  }

  render(){

    return (

        <Container component="main" maxWidth="xs">
          <CssBaseline />
          <div className={"blank"}>
            <Typography component="h1" variant="h5">
              {"  Register Account  "}
            </Typography>
            <form className={"blank"} noValidate>
              <Grid container spacing={2}>
                <Grid item xs={12} sm={6}>
                  <TextField
                    autoComplete="fname"
                    name="firstName"
                    value={this.state.firstName}
                    onChange={(e)=>{this.setState({firstName: e.target.value})}}
                    variant="outlined"
                    required
                    fullWidth
                    id="firstName"
                    label="First Name"
                    autoFocus
                  />
                </Grid>
                <Grid item xs={12} sm={6}>
                  <TextField
                    variant="outlined"
                    required
                    fullWidth
                    id="lastName"
                    value={this.state.lastName}
                    onChange={(e)=>{this.setState({lastName: e.target.value})}}
                    label="Last Name"
                    name="lastName"
                    autoComplete="lname"
                  />
                </Grid>
                <Grid item xs={12}>
                  <TextField
                    variant="outlined"
                    required
                    fullWidth
                    id="email"
                    value={this.state.email}
                    onChange={(e)=>{this.setState({email: e.target.value})}}
                    label="Email Address"
                    name="email"
                    autoComplete="email"
                  />
                </Grid>
                <Grid item xs={12}>
                  <TextField
                    variant="outlined"
                    required
                    fullWidth
                    name="password"
                    value={this.state.password}
                    onChange={(e)=>{this.setState({password: e.target.value})}}
                    label="Password"
                    type="password"
                    id="password"
                    autoComplete="current-password"
                  />
                </Grid>
                
              </Grid>
              <Button
                onClick={this.defaultPreventer}              
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={"blank"}
              >
                Sign Up
              </Button>

              <Grid container justify="flex-end">
                <Grid item>
                  <Link onClick={()=>this.props.toLogIn()} href="#" variant="body2">
                    {"Sign in"}
                  </Link>
                </Grid>
              </Grid>

            </form>
          </div>
          
        </Container>
      );
  }
}