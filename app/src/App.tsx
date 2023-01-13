import axios from 'axios'
import React from 'react'
import logo from './logo.svg'
import config from './config.json'
import './App.css'

function App (): JSX.Element {
  const [questions, setQuestions] = React.useState(null)

  React.useEffect(() => {
    const fetchData = async (): Promise<void> => {
      const results = await axios.get(`${config.apiBaseUrl}/questions`)
      console.log('Got data succesfully:', results.data)
      setQuestions(results.data)

      console.log(questions)
    }

    fetchData().catch((error) => { console.error('Error in in fetching data:', error) })
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
  )
}

export default App
