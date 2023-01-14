import React from 'react'
import { render } from '@testing-library/react'
import App from './App'

test('renders layout in app', () => {
  render(<App />)
  const header = document.querySelector('#app-header')
  const content = document.querySelector('#app-content')
  const footer = document.querySelector('#app-footer')

  expect(header).toBeInTheDocument()
  expect(content).toBeInTheDocument()
  expect(footer).toBeInTheDocument()
})
