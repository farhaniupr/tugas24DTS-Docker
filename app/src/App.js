import logo from './logo.svg';
import './App.css';

import React, {useEffect} from 'react';
import axios from 'axios';

import ReactDOM from 'react-dom';
import { Router, Route, Link, browserHistory, IndexRoute } from 'react-router'

class App extends React.Component {
  render() {
     return (
        <div>
           <ul>
           <li>Login</li>
           <li>About</li>
           <li>Contact</li>
           </ul>
           {this.props.children}
        </div>
     )
  }
}

export default App;

class Login extends React.Component {
  state = {
    id_account: '',
    password : '',
  }

  handleChange = event => {
    this.setState({ 
      id_account: event.target.value, 
      //password  : event.target.value,
    });
  }

  handleChange2 = event => {
    this.setState({  
      password  : event.target.value,
    });
  }

  handleSubmit = event => {
    event.preventDefault();

    /*const Account = {
      id_account: this.state.id_account,
      password: this.state.password
    };*/

    axios.post(`http://localhost:8084/guest/login`, {id_account: this.state.id_account, password: this.state.password})
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
            Id Account:
            <input type="text" name="id_account" onChange={this.handleChange} />
          </label>
          <label>
            Password:
            <input type="text" name="password" onChange={this.handleChange2} />
          </label>
          <button type="submit">Add</button>
        </form>
      </div>
    )
  }
}

//export default App;
