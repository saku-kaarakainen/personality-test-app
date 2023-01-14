import React from 'react'
import MagicUser from '@testing-library/user-event'
import { act, render, screen } from '@testing-library/react'
import LandingPage from './LandingPage'

const mockedUsedNavigate = jest.fn()
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockedUsedNavigate
}))

test('renders layout in app', () => {
  render(<LandingPage />)

  const submitButton = screen.getByRole('button')

  act(() => {
    MagicUser.click(submitButton)
  })

  expect(mockedUsedNavigate).toHaveBeenCalledWith('Questions')
  expect(mockedUsedNavigate).toHaveBeenCalledTimes(1)
})
