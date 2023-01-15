import React from 'react'
import axios from 'axios'
import config from './../config.json'
import Question, { QuestionSet } from '../components/Question'

const Questions: React.FC = () => {
  const [questions, setQuestions] = React.useState(new Array<QuestionSet>())
  const [currentStep] = React.useState(1)
  const [maxSteps, setMaxSteps] = React.useState(1)

  React.useEffect(() => {
    const fetchData = async (): Promise<void> => {
      const results = await axios.get(`${config.apiBaseUrl}/questions`)
      setQuestions(results.data)
      setMaxSteps(questions.length)
    }

    fetchData().catch((error) => { console.error('Error in in fetching data:', error) })
  }, [])

  /*
    TODO:
     - next question
     - previous question
     - tests
  */

  if (questions.length === 0) {
    // TODO: Show loading spinner, when the results are not available
    return <></>
  }

  return <Question
    questions={questions}
    currentStep={currentStep}
    maxSteps={maxSteps} />
}

export default Questions
