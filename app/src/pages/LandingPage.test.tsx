import React from 'react'
import { render } from '@testing-library/react'
import LandingPage from './LandingPage'

test('renders layout in app', () => {
  render(<LandingPage />)
})
