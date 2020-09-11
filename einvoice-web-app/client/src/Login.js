import React, { Component } from 'react'
import './App.css'

class Login extends Component {
  render() {
    return (
      <div className="App container">
        <header className="App-header row">
          <h1 className='col'>E-invoice</h1>
        </header>
        <div className='row justify-content-center'>
          <button className='btn btn-primary col-sm-2' onClick={this.props.login}>Login</button>
        </div>
      </div>
    )
  }
}

export default Login;