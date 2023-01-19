// import axios from 'axios'
import React from 'react'
import axios from 'axios'
import { useLocation, useNavigate } from 'react-router-dom'
import { AnswerSet } from './Questions'
import config from './../config.json'
import './Result.css'
import { Button } from 'antd'

interface ApiResult {
  id: string
  label: string
  description_paragraphs: string[]
}

const Result: React.FC = () => {
  const navigate = useNavigate()
  const routeChange = (): void => {
    navigate('/')
  }
  const location = useLocation()
  const { state } = location
  const { answers } = state as { answers: AnswerSet[] }
  const params = answers
    .map(a => `q[${a.question_id}]=${a.answer}`)
    .join('&')
  const [apiResult, setApiResult] = React.useState<ApiResult>({
    id: '',
    label: '',
    description_paragraphs: []
  })

  React.useEffect(() => {
    const fetchData = async (): Promise<void> => {
      const response = await axios.get(`${config.apiBaseUrl}/result/calculate?${params}`)

      setApiResult(response.data)
    }

    fetchData()
      .catch((error) => {
        console.error('Error in in fetching data:', error)
      })
  })

  return <div className='result-container'>
    <h3>You Result</h3>
    <h4>{apiResult.label}</h4>
    <div className='result-text'>
      <p>{apiResult.description_paragraphs.join('<br />')}</p>
    </div>
    <div className='go-home-area'>
      <Button type="primary" onClick={routeChange}>
        <span id="home-page">Home Page</span>
      </Button>
    </div>
  </div>
}

export default Result
