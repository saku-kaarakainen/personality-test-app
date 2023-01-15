import React from 'react'
import { Radio, Space } from 'antd'

export interface QuestionSet {
  id: string
  question_text: string
  question_description: string
  answers: Array<{
    id: string
    answer_label: string
  }>
}

interface QuestionProps {
  questions: QuestionSet[]
  currentStep: number
}

const Question: React.FC<QuestionProps> = (props: QuestionProps) => {
  console.log('currentStep:', props.currentStep)

  const currentQuestion = props.questions[props.currentStep]
  console.log('currentQuestion:', currentQuestion)

  /*
    TODO:
     - current step breadcrump
     - pass radiobutton value to parent
  */
  return <>
    <h3 className="question-text">{currentQuestion.question_text}</h3>
    <div className="question-description">
      <span>{currentQuestion.question_description}</span>
    </div>
    <Radio.Group>
      <Space direction='vertical'>
        {currentQuestion.answers.map(a => {
          return <Radio value={a.id} key={a.id}>{a.answer_label}</Radio>
        })}
      </Space>
    </Radio.Group>
  </>
}

export default Question
