import './App.css';

import React, {useEffect} from 'react';
import axios from 'axios';

export default class doWithDraw extends React.Component {
  state = {
    ID: null,
    transaction_type: '',
    transaction_description: '',
    sender : '',
    amount: null,
    recipient: '',
    timestamp:'',
  }

  handleChange = event => {
    this.setState({ 
      ID: event.target.value, 
    });
  }

  handleChange2 = event => {
    this.setState({  
        transaction_type  : event.target.value,
    });
  }

  handleChange3 = event => {
    this.setState({ 
        transaction_description: event.target.value, 
    });
  }

  handleChange4 = event => {
    this.setState({  
      sender  : event.target.value,
    });
  }

  handleChange5 = event => {
    this.setState({  
      amount  : event.target.value,
    });
  }

  handleChange6 = event => {
    this.setState({  
        recipient  : event.target.value,
    });
  }
  handleChange7 = event => {
    this.setState({  
        timestamp: event.target.value,
    });
  }

  handleSubmit = event => {
    event.preventDefault();    

    axios.post(`http://localhost:8084/guest/withdraw`, {ID: parseInt(this.state.ID,10), transaction_type: this.state.transaction_type, transaction_description: this.state.transaction_description, sender: this.state.sender, amount: parseInt(this.state.amount), recipient: this.state.recipient, timestamp:this.state.timestamp})
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
          <p></p>
          <label>
            transaction type:
            <input type="text" name="transaction_type" onChange={this.handleChange2} />
          </label>   
          <p></p>
          <label>
            transaction description:
            <input type="text" name="transaction_description" onChange={this.handleChange3} />
          </label>
          <p></p>
          <label>
            sender:
            <input type="text" name="sender" onChange={this.handleChange4} />
          </label>
          <p></p>
          <label>
            Amount:
            <input type="number" name="amount" onChange={this.handleChange5} />
          </label>
          <p></p>
          <label>
            recipient:
            <input type="text" name="recipient" onChange={this.handleChange6} />
          </label>
          <p></p>
          <label>
            timestamp:
            <input type="text" name="timestamp" onChange={this.handleChange7} />
          </label>
          <p></p>
          <button type="submit">Withdraw</button>
        </form>
      </div>
    )
  }
}

