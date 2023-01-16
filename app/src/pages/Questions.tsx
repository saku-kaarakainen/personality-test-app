import React from 'react'
import axios from 'axios'
import { Button, Row, Col } from 'antd'
import config from './../config.json'
import Question, { QuestionSet } from '../components/Question'
import './Questions.css'

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

  return (<>
    <Row>
      <Col span={24} className='col-question'>
        <Question
          questions={questions}
          currentStep={currentStep}
          maxSteps={maxSteps} />
      </Col>
    </Row>
    <Row>
      <Col span={12} className='col-prev'><Button>Previous</Button></Col>
      <Col span={12} className='col-next'><Button type="primary">Next</Button></Col>
    </Row>
  </>)
}

export default Questions
