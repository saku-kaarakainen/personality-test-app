import React from 'react'
import MagicUser from '@testing-library/user-event'
import { act, render, screen } from '@testing-library/react'
import Questions from './Questions'

const mockedUsedNavigate = jest.fn()
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockedUsedNavigate
}))

beforeEach(() => {

})

test('renders layout in app', () => {
  const component = render(<Questions />)
  const prevButton = document.querySelector('.col-prev button')!
  const nextButton = document.querySelector('.col-next button')!


  act(() => {
    MagicUser.click(prevButton)
  })
})
