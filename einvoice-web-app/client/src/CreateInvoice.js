import React, { Component } from "react";
import { defaultUbl, defaultD16b } from "./default";

class CreateInvoice extends Component {
  constructor(props) {
    super(props);

    this.state = {
      getInvoice: null,
      getInvoiceId: "",
      postInvoiceSender: "",
      postInvoiceReceiver: "",
      format: "json",
      xmlInputValue: "",
    }

    this.handleInputChange = this.handleInputChange.bind(this);
    this.submitJsonInvoice = this.submitJsonInvoice.bind(this);
    this.selectFormat = this.selectFormat.bind(this);
    this.updateXmlInputValue = this.updateXmlInputValue.bind(this);
    this.submitXmlInvoice = this.submitXmlInvoice.bind(this);
  }

  handleInputChange(event) {
    const target = event.target;
    const name = target.name;
    this.setState({
      [name]: target.value
    });
  }

  selectFormat(event) {
    this.setState({ format: event.target.name });
    if(event.target.name === 'ubl') {
      this.setState({ xmlInputValue: defaultUbl });
    } else if(event.target.name === 'd16b') {
      this.setState({ xmlInputValue: defaultD16b });
    }
  }

  updateXmlInputValue(event) {
    this.setState({ xmlInputValue: event.target.value });
  }

  submitJsonInvoice() {
    fetch(this.props.apiUrl + '/api/invoice/json', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': this.props.user.token
      },
      body: JSON.stringify({
        sender: this.state.postInvoiceSender,
        receiver: this.state.postInvoiceReceiver
      })
    })
      .then( response => response.json())
      .then( data => {
        this.props.addInvoice(data);
        this.setState({postInvoiceSender: "", postInvoiceReceiver: ""})
      });
  }

  submitXmlInvoice() {
    fetch(this.props.apiUrl + '/api/invoice/' + this.state.format, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/xml',
        'Authorization': this.props.user.token
      },
      body: this.state.xmlInputValue
    })
      .then( response => response.json())
      .then( data => {
        this.props.addInvoice(data);
        this.setState({postInvoiceSender: "", postInvoiceReceiver: ""})
      });
  }

  render() {
    let { format } = this.state;
    let form = null;
    if(format === 'json') {
      form = <div>
        <div className="row">
          <div className="col"><p>
            Sender
            <input type="text" name="postInvoiceSender" value={this.state.postInvoiceSender}
                   onChange={this.handleInputChange}/>
          </p>
          </div>
        </div>
        <div className="row">
          <div className="col">
            <p>
              Receiver
              <input type="text" name="postInvoiceReceiver" value={this.state.postInvoiceReceiver}
                     onChange={this.handleInputChange}/>
            </p>
            <button className='btn btn-primary' onClick={this.submitJsonInvoice}>Submit</button>
          </div>
        </div>
      </div>
    } else {
      form = (
        <div>
          <div className='row justify-content-center'>
            <textarea className='col' name="xml" cols="50" rows="15" value={this.state.xmlInputValue} onChange={this.updateXmlInputValue}/>
          </div>
          <div className='row justify-content-center'>
            <button className='btn btn-primary' onClick={this.submitXmlInvoice}>Submit</button>
          </div>
        </div>
      )
    }

    return (
      <div className="container">
        <div className="row justify-content-center">
          <button className='btn btn-primary col-sm-2' name="json" onClick={this.selectFormat}>Json</button>
          <button className='btn btn-primary col-sm-2' name="ubl" onClick={this.selectFormat}>UBL2.1</button>
          <button className='btn btn-primary col-sm-2' name="d16b" onClick={this.selectFormat}>D16B</button>
        </div>
        <p className='row justify-content-center'>Format: { format }</p>
        { form }
      </div>
    )
  }
}

export default CreateInvoice;