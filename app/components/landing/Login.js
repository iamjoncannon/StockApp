import React from 'react';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';

export default class SignIn extends React.Component {

  constructor(props) {
    super(props);
    this.state = { 
      email: "",
      password: ""
    }
  }

  defaultPreventer = (evt) => {

    evt.preventDefault()
    evt.stopPropagation()
    this.props.handleLogIn(this.state.email, this.state.password)
  }

  render() { 

      return (

        <Container component="main" maxWidth="xs">
          <CssBaseline />
          <div className={"blank" } >
            <Typography component="h1" variant="h5">
              Log Into Trading Account
            </Typography>
            <form className={"blank" } noValidate>
              <TextField
                variant="outlined"
                value={this.state.email}
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                onChange={(e)=>{this.setState({email: e.target.value})}}
                autoFocus
              />
              <TextField
                variant="outlined"
                value={this.state.password}
                onChange={(e)=>{this.setState({password: e.target.value})}}
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"            
                autoComplete="current-password"
              />
              <Button
                onClick={this.defaultPreventer}
                fullWidth
                variant="contained"
                color="primary"
                className={"blank"}
              >
                Sign In
              </Button>
              <Grid container>
                  <Link onClick={this.props.toSignUp} href="#" variant="body2">
                    {"Register Account"}
                  </Link>            
              </Grid>
            </form>
          </div>
        </Container>
      );
  }
}

// source: https://github.com/mui-org/material-ui/blob/master/docs/src/pages/getting-started/page-layout-examples/sign-in/SignIn.js
