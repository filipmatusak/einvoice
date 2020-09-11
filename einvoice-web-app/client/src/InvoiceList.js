import React, { Component } from 'react'
import './App.css'

class InvoiceList extends Component {
  render() {
    let { invoices } = this.props;

    let rows = [];
    if(invoices) {
      rows = invoices.map((invoice, i) => {
        return <tr key={i}>
          <th scope="row"><p>{i+1}</p></th>
          <td><p>{invoice.id}</p></td>
          <td><p>{invoice.sender}</p></td>
          <td><p>{invoice.receiver}</p></td>
          <td><p>{invoice.price}</p></td>
          <td><p>{invoice.format}</p></td>
          <td><p className='link' onClick={() => this.props.viewInvoice(invoice.id)}>view</p></td>
        </tr>
      });
    }

    return (
      <div className="container">
        <h2>All invoices</h2>
        <table className="table table-striped table-borderless table-hover">
          <thead>
          <tr>
            <th>#</th>
            <th>ID</th>
            <th>Sender</th>
            <th>Receiver</th>
            <th>Price</th>
            <th>Format</th>
            <th></th>
          </tr>
          </thead>
          <tbody>
          {rows}
          </tbody>
        </table>
      </div>
    )
  }
}

export default InvoiceList;