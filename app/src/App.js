import logo from './logo.svg';
import './App.css';

import React, {useEffect} from 'react';
import axios from 'axios';

import Createaccount from './Createaccount';
import Login from './Login';
import Deposit from './Deposit';
import Withdraw from './Withdraw';
import Transfer from './Transfer';
import GetAcc from './GetAcc';

import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

export default function BasicExample() {
  return (
    <Router>
      <div>
        <ul>
          <li>
            <Link to="/Login">Login</Link>
          </li>
          <li>
            <Link to="/Createaccount">Create Account</Link>
          </li>
          <li>
            <Link to="/Deposit">Deposit</Link>
          </li>
          <li>
            <Link to="/Withdraw">Withdraw</Link>
          </li>
          <li>
            <Link to="/Transfer">Transfer</Link>
          </li>
          <li>
            <Link to="/GetAcc">Mutasi</Link>
          </li>
        </ul>

        <hr />

        {/*
          A <Switch> looks through all its children <Route>
          elements and renders the first one whose path
          matches the current URL. Use a <Switch> any time
          you have multiple routes, but you want only one
          of them to render at a time
        */}
        <Switch>
          <Route exact path="/Login">
            <Login />
            </Route>
          <Route path="/Createaccount">
            <Createaccount/>
          </Route>
          <Route path="/Deposit">
            <Deposit />
          </Route>
          <Route path="/Withdraw">
            <Withdraw/>
          </Route>
          <Route path="/Transfer">
            <Transfer/>
          </Route>
          <Route path="/GetAcc">
            <GetAcc/>
          </Route>
        </Switch>
      </div>
    </Router>
  );
}
