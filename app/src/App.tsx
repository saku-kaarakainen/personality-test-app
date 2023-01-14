import axios from 'axios'
import React from 'react'
import { Layout } from 'antd'
import config from './config.json'
import './App.css'

const { Header, Footer, Content } = Layout

// TODO:
// use breadcrumb
// https://ant.design/components/breadcrumb

// Radiobutton

const App: React.FC = () => {
  const [questions, setQuestions] = React.useState(null)

  React.useEffect(() => {
    const fetchData = async (): Promise<void> => {
      const results = await axios.get(`${config.apiBaseUrl}/questions`)
      console.log('Got data succesfully:', results.data)
      setQuestions(results.data)

      console.log(questions)
    }

    fetchData().catch((error) => { console.error('Error in in fetching data:', error) })
  })

  return (
    <>
      {/* Only one layout, that is mobile-friendly */}
      <Layout>
        <Header id="app-header"><h1>Personality Test App</h1></Header>
        <Content id="app-content">Content</Content>
        <Footer id="app-footer"></Footer>
      </Layout>
    </>
  )
}

export default App
