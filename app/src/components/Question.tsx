import React from 'react'
import { Breadcrumb, Radio, Space } from 'antd'
import './Question.css'

export interface QuestionSet {
  id: string
  question_text: string
  question_description: string
  answers: Array<{
    id: string
    answer_label: string
  }>
}

const Question: React.FC< {
  questions: QuestionSet[]
  currentStep: number
  maxSteps: number
}> = ({
  currentStep,
  questions,
  maxSteps
}) => {
  const currentQuestion = questions[currentStep]

  /*
    TODO:
     - current step breadcrump
     - pass radiobutton value to parent
  */
  return <>
    <Breadcrumb>
      <Breadcrumb.Item className="question-breadcrumb">Question {currentStep + 1}/{maxSteps}</Breadcrumb.Item>
    </Breadcrumb>

    <h3 className="question-text">{currentQuestion.question_text}</h3>
    <div className="question-description">
      <span>{currentQuestion.question_description}</span>
    </div>
    <Radio.Group className='question-group'>
      <Space direction='vertical'>
        {currentQuestion.answers.map(a => {
          return <Radio value={a.id} key={`key-${a.id}`}>
            <div className='question-item'>
              <span className="question-item-text">{a.answer_label}</span>
            </div>
          </Radio>
        })}
      </Space>
    </Radio.Group>
  </>
}

export default Question
