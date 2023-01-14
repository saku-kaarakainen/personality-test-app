import React from 'react'
import { render } from '@testing-library/react'
import LandingPage from './LandingPage'

const mockedUsedNavigate = jest.fn();
jest.mock('react-router-dom', () => ({
   ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockedUsedNavigate,
}));

test('renders layout in app', () => {
  render(<LandingPage />)
})
