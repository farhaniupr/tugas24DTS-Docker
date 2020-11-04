import React, {useEffect} from 'react';
import axios from 'axios';
import './App.css';


function App() {
  const postData = () => {
    axios.post('http://localhost:8084/guest/login', {id_account: '1xsDASqwe', password : 'Farhan'})
      .then(response => console.log(response))
      .catch(error => console.log(error))
  }
  return (
    <div className="App">
      <header className="App-header">
      <label for="IdAccount">Id Account</label>
        <input name="IdAccount" id="IdAccount"/>
        
        <label for="Password">Password</label>
        <input name='Password' id="Password"/>
        
        <button onClick={postData}>Login</button>
      </header>
    </div>
  );  
}

export default App;
