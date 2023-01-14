import React from 'react'
import { render, fireEvent } from '@testing-library/react'
import App from './App'
import { MemoryRouter } from 'react-router-dom';
import { act } from 'react-dom/test-utils';

test('renders layout in app', () => {
  render(<App />)
  const header = document.querySelector('#app-header')
  const content = document.querySelector('#app-content')
  const footer = document.querySelector('#app-footer')

  expect(header).toBeInTheDocument()
  expect(content).toBeInTheDocument()
  expect(footer).toBeInTheDocument()
})

test('title link navigates to landing page', () => {
  render(<App />)

  const goHomeLink = document.querySelector("#title-link");

  act(() => {
    goHomeLink?.dispatchEvent(new MouseEvent("click", { bubbles: true }))
  })
  
  expect(document.body.textContent).toContain("Take a test")
})