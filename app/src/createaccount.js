import './App.css';

import React, {useEffect} from 'react';
import axios from 'axios';

export default class doCreate extends React.Component {
  state = {
    ID: '',
    id_account: '',
    name : '',
    password : '',
    account_number: '',
    saldo: '',
  }

  handleChange = event => {
    this.setState({ 
      ID: event.target.value, 
    });
  }

  handleChange2 = event => {
    this.setState({  
      id_account  : event.target.value,
    });
  }

  handleChange3 = event => {
    this.setState({ 
      name: event.target.value, 
    });
  }

  handleChange4 = event => {
    this.setState({  
      password  : event.target.value,
    });
  }

  handleChange5 = event => {
    this.setState({  
      account_number  : event.target.value,
    });
  }

  handleChange6 = event => {
    this.setState({  
      saldo  : event.target.value,
    });
  }

  handleSubmit = event => {
    event.preventDefault();


    axios.post(`http://localhost:8084/guest/create`, {ID: this.state.ID, id_account: this.state.id_account, name : this.state.name, password: this.state.password, account_number: this.state.account_number, saldo:this.state.saldo})
      .then(res => {
        console.log(res);
        console.log(res.data);
      })
  }

  render() {
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
        <label>
            Id :
            <input type="number" name="ID" onChange={this.handleChange} />
          </label>
          <label>
            ID Account:
            <input type="text" name="id_account" onChange={this.handleChange2} />
          </label>   
          <label>
            Nama:
            <input type="text" name="account_number" onChange={this.handleChange3} />
          </label>
          <label>
            Password:
            <input type="text" name="password" onChange={this.handleChange4} />
          </label>
          <label>
            Account Number:
            <input type="number" name="account_number" onChange={this.handleChange5} />
          </label>
          <label>
            Saldo:
            <input type="number" name="saldo" onChange={this.handleChange6} />
          </label>
          <button type="submit">Add</button>
        </form>
      </div>
    )
  }
}

