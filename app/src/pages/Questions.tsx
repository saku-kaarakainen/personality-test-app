import React from 'react'
import axios from 'axios'
import config from './../config.json'
import Question, { QuestionSet } from '../components/Question'

const Questions: React.FC = () => {
  const [questions, setQuestions] = React.useState(new Array<QuestionSet>())
  const [currentStep, setCurrentStep] = React.useState(0)

  const notTrue = false
  if (notTrue) {
    setCurrentStep(0)
  }

  React.useEffect(() => {
    const fetchData = async (): Promise<void> => {
      const results = await axios.get(`${config.apiBaseUrl}/questions`)
      console.log('Got data succesfully:', results.data)

      const qsts: QuestionSet[] = results.data
      setQuestions(qsts)

      console.log(questions)
    }

    fetchData().catch((error) => { console.error('Error in in fetching data:', error) })
  })

  /*
    TODO:
     - next question
     - previous question
     - tests
  */
  return <Question questions={questions} currentStep={currentStep} />
}

export default Questions
