import axios from 'axios';
import React from 'react';
import logo from './logo.svg';
import './App.css';

const baseURL = "http://localhost:8080"

function App() {
  const [questions, setQuestions] = React.useState(null);

  React.useEffect(() => {
    const fetchData = async () => {
      const results = await axios.get(`${baseURL}/questions`)
      console.log("Got data succesfully:", results)
    }

    fetchData()
      .catch((error) => {
        console.error("Error in in fetching data:", error)
      })
   })

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
