import React from 'react'

export type QuestionSet = {
  id: string,
  question_text: string,
  question_description: string,
  answers: {
    id: string,
    answer_label: string
  }
}

type QuestionProps = {
  questions: QuestionSet[],
  currentStep: number
}

const Question: React.FC<QuestionProps> = (props: QuestionProps) => {
  console.log("currentStep:", props.currentStep)
  
  const currentQuestion = props.questions[props.currentStep]
  console.log("currentQuestion:", currentQuestion)

  /*
    TODO:
     - current step breadcrump
     - questions as radiobuttons
     - pass radiobutton value to parent
  */
  return <>
    {currentQuestion}
  </>
}

export default Question
