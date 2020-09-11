import React, { Component } from 'react'
import './App.css'

class InvoiceView extends Component {
  render() {
    return (
      <div className="container">
        <div className='row justify-content-center'>
          <button className='btn btn-primary col-sm-2' onClick={this.props.close}>Close</button>
        </div>
        <div className='row justify-content-center'>
          <textarea rows="40" cols="100">{ this.props.invoice }</textarea>
        </div>
      </div>
    )
  }
}

export default InvoiceView;