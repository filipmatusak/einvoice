import React, {Component} from 'react';
import './App.css';
import CreateInvoice from './CreateInvoice';
import InvoiceList from './InvoiceList';
import Login from "./Login";
import InvoiceView from "./InvoiceView";

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      apiUrl: null,
      invoices: [],
      createInvoice: false,
      user: null,
      viewedInvoice: null
    }

    this.addInvoice = this.addInvoice.bind(this);
    this.createInvoice = this.createInvoice.bind(this);
    this.login = this.login.bind(this);
    this.logout = this.logout.bind(this);
    this.closeInvoice = this.closeInvoice.bind(this);
    this.viewInvoice = this.viewInvoice.bind(this);
  }

  addInvoice(invoice) {
    let invoices = this.state.invoices;
    invoices.push(invoice);
    this.setState({invoices});
  }

  login() {
    fetch(this.state.apiUrl + "/login")
      .then( response => response.json())
      .then( user => {
        this.setState({ user });
        localStorage.setItem('user', JSON.stringify(user));
        this.getInvoices();
      });
  }

  logout() {
    fetch(this.state.apiUrl + "/logout", {
      headers: {
        Authorization: this.state.user.token
      }
    });
    this.setState({ user: null });
    localStorage.removeItem('user');
  }

  getInvoices() {
    fetch(this.state.apiUrl + '/api/invoices', {
      headers: {
        Authorization: this.state.user.token
      }
    })
      .then( response => response.json())
      .then( data => {
        this.setState({invoices: data});
      });
  }

  createInvoice(){
    this.setState({ createInvoice: !this.state.createInvoice});
  }

  closeInvoice(){
    this.setState({ viewedInvoice: null });
  }

  viewInvoice(id){
    fetch(this.state.apiUrl + '/api/invoice/full/' + id , {
      headers: {
        Authorization: this.state.user.token
      }
    }).then( response => response.text())
      .then( invoice => this.setState({viewedInvoice: invoice}));
  }

  componentDidMount() {
    fetch('/api/url')
      .then( response => response.text())
      .then( url => {
        this.setState({apiUrl: url});
        let userString = localStorage.getItem('user');
        let user = userString && JSON.parse(userString);
        if(user) {
          fetch(this.state.apiUrl + '/me', {
            headers: {
              Authorization: user.token
            }
          }).then( response => {
            if(response.status === 401) {
              this.setState({ user: null })
            } else {
              response.json().then(user => {
                this.setState({ user });
                this.getInvoices();
              });
            }
          });
        }
      });
  }

  render() {
    let { apiUrl, user, viewedInvoice } = this.state;

    if(apiUrl) {
      if(user) {
        return (
          <div className="App container">
            <header className="App-header row">
              <h1 className='col'>E-invoice</h1>
            </header>
            <div className='row justify-content-center'>
              <p className='col-sm-1'>User id:</p>
              <p className='col-sm-1'>{user.id}</p>
              <button className='btn btn-primary' onClick={this.logout}>Logout</button>
            </div>
            { !viewedInvoice &&
            <div>
              <div className='row justify-content-center'>
                <button className='btn btn-primary' onClick={this.createInvoice}>Create invoice</button>
              </div>
              {this.state.createInvoice && <CreateInvoice apiUrl={apiUrl} addInvoice={this.addInvoice} user={user}/>}
              <InvoiceList invoices={this.state.invoices} apiUrl={apiUrl} user={user} viewInvoice={this.viewInvoice}/>
            </div>
            }
            { viewedInvoice && <InvoiceView close={this.closeInvoice} invoice={viewedInvoice}/> }
          </div>
        );
      } else return <Login login={this.login}/>
    } else return <div></div>;
  }
}

export default App;
