import React from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'
import { Button, Row, Col } from 'antd'
import config from './../config.json'
import Question, { QuestionSet } from '../components/Question'
import './Questions.css'

const Questions: React.FC = () => {
  const navigate = useNavigate()
  const [questions, setQuestions] = React.useState(new Array<QuestionSet>())
  const [currentStep, setCurrentStep] = React.useState(0)

  React.useEffect(() => {
    const fetchData = async (): Promise<void> => {
      const results = await axios.get(`${config.apiBaseUrl}/questions`)

      setQuestions(results.data)
    }

    fetchData().catch((error) => { console.error('Error in in fetching data:', error) })
  }, [questions])

  if (questions.length === 0) {
    // TODO: Show loading spinner, when the results are not available
    return <></>
  }

  const onPrevClick = (): void => {
    if (currentStep < 1) {
      // First step, reset to default
      setCurrentStep(0)
      navigate('/')
      return
    }

    setCurrentStep(currentStep - 1)
  }

  const onNextClick = (): void => {
    const lastIndex = questions.length - 1
    if (currentStep < lastIndex) {
      setCurrentStep(currentStep + 1)
    } else {
      navigate('/Result')
    }
  }

  return (<>
    <Row>
      <Col span={24} className='col-question'>
        <Question
          questions={questions}
          currentStep={currentStep}
          maxSteps={questions.length} />
      </Col>
    </Row>
    <Row>
      <Col span={12} className='col-prev' onClick={onPrevClick}>
        <Button>{currentStep === 0 ? 'Home page' : 'Previous'}</Button>
      </Col>
      <Col span={12} className='col-next' onClick={onNextClick}>
        <Button type="primary">{currentStep === questions.length - 1 ? 'See results' : 'Next'}</Button>
      </Col>
    </Row>
  </>)
}

export default Questions
