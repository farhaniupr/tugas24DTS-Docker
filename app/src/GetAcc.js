import './App.css';

import React, {useEffect} from 'react';
import axios from 'axios';

export default class GetAcc extends React.Component {
  state = {
    id_account: '',
  }

  handleChange = event => {
    this.setState({ 
      id_account: event.target.value, 
    });
  }


  handleSubmit = event => {
    event.preventDefault();    

    axios.post(`http://localhost:8084/guest/detailac`, {id_account: this.state.id_account})
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
            Id Account :
            <input type="text" name="id_account" onChange={this.handleChange} />
          </label>
          <p></p>

          <button type="submit">Tampilkan</button>
        </form>
      </div>
    )
  }
}

