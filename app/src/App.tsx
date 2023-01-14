import React from 'react'
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom'
import { Layout } from 'antd'
import LandingPage from './pages/LandingPage'
import Questions from './pages/Questions'
import Result from './pages/Result'
import NotFound from './pages/NotFound'
import './App.css'

const { Header, Footer, Content } = Layout

// TODO:
// use breadcrumb
// https://ant.design/components/breadcrumb

// Radiobutton

const App: React.FC = () => {
  return (
    <>
      <BrowserRouter>
        {/* Only one layout, that is mobile-friendly */}
        <Layout>
          <Header id="app-header">
            <h1>
              <Link id="title-link" to="/">Personality Test App</Link>
            </h1>
          </Header>
          <Content id="app-content">
            <h2>Are you an introvert or an extrovert?</h2>
            <Routes>
              <Route path="/" element={<LandingPage />} />
              <Route path="questions" element={<Questions />} />
              <Route path="result" element={<Result />} />
              <Route path="*" element={<NotFound />} />
            </Routes>
          </Content>
          <Footer id="app-footer">

          </Footer>
        </Layout>
      </BrowserRouter>
    </>
  )
}

export default App
