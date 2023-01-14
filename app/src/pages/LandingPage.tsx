import React from 'react'
import { Button, Space } from 'antd'
import { useNavigate } from 'react-router-dom'

const LandingPage: React.FC = () => {
  const navigate = useNavigate()
  const routeChange = (): void => {
    navigate('Questions')
  }

  return <Space>
    <Button type="primary" onClick={routeChange}>
      <span id="start-test">Take a test</span>
      </Button>
  </Space>
}

export default LandingPage
