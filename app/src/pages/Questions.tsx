import React from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'
import { Button, Row, Col, RadioChangeEvent } from 'antd'
import config from './../config.json'
import Question, { QuestionSet } from '../components/Question'
import './Questions.css'

interface AnswerSet {
  question_id: string
  answer: string
}

const Questions: React.FC = () => {
  const navigate = useNavigate()
  const [questions, setQuestions] = React.useState(new Array<QuestionSet>())
  const [answers, setAnswers] = React.useState(new Array<AnswerSet>())
  const [currentStep, setCurrentStep] = React.useState(0)
  const [isProceedDisabled, setIsProceedDisabled] = React.useState(false)

  React.useEffect(() => {
    const fetchData = async (): Promise<void> => {
      const results = await axios.get(`${config.apiBaseUrl}/questions`)

      setQuestions(results.data)

      // initializes the answers
      if (answers === undefined || answers.length === 0) {
        setAnswers(
          (results.data as QuestionSet[])
            .map(q => {
              return {
                question_id: q.id,
                answer: ''
              }
            })
        )
      }
    }

    fetchData().catch((error) => { console.error('Error in in fetching data:', error) })
  }, [questions])

  // Logic to check if navigation to next page/step is allowed
  React.useEffect(() => {
    if (answers === undefined || answers.length === 0) {
      setIsProceedDisabled(true)
      return
    }

    const current = answers.filter(a => a.question_id === getCurrentQuestion().id)[0]
    setIsProceedDisabled(current.answer === undefined || current.answer === '')
  }, [
    // Update based on questions, because it just seems to work the best
    questions
  ])

  if (questions.length === 0) {
    // TODO: Show loading spinner, when the results are not available
    return <></>
  }

  const getCurrentQuestion = (): QuestionSet => questions[currentStep]

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

  const onValueChange = (e: RadioChangeEvent): void => {
    const newAnswerSet = answers
      .map(a => a.question_id === getCurrentQuestion().id
        ? { ...a, answer: e.target.value }
        : a)
    setAnswers(newAnswerSet)
  }

  return (<>
    <Row>
      <Col span={24} className='col-question'>
        <Question
          currentQuestion={getCurrentQuestion()}
          currentStep={currentStep}
          maxSteps={questions.length}
          onRadioChange={onValueChange} />
      </Col>
    </Row>
    <Row>
      <Col span={12} className='col-prev' onClick={onPrevClick}>
        <Button>{currentStep === 0 ? 'Home page' : 'Previous'}</Button>
      </Col>
      <Col span={12} className='col-next' onClick={onNextClick}>
        <Button type="primary" disabled={isProceedDisabled} >{currentStep === questions.length - 1 ? 'See results' : 'Next'}</Button>
      </Col>
    </Row>
  </>)
}

export default Questions
