import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import Createaccount from './Createaccount';
import doCreate from './Createaccount';
import Login from './Login';
import Deposit from './Deposit';
import Withdraw from './Withdraw';
import Transfer from './Transfer';
import GetAcc from './GetAcc';

ReactDOM.render(
  <React.StrictMode>
    <App />
    <Login/> 
    <Createaccount/>
    <Deposit/>
    <Withdraw/>
    <Transfer/>
    <GetAcc/>
  </React.StrictMode>,
  document.getElementById('root')
);

reportWebVitals();
