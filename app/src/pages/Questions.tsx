import React from 'react'
import axios from 'axios'
import config from './../config.json'

const Questions: React.FC = () => {
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

  return <p>Page Results - Questions</p>
}

export default Questions
