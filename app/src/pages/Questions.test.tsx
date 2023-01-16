import React from 'react'
import MagicUser from '@testing-library/user-event'
import { act, render, screen, waitFor } from '@testing-library/react'
import Questions from './Questions'
import axios from 'axios'

const mockedUsedNavigate = jest.fn()
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockedUsedNavigate
}))

jest.mock('axios');
const mockedAxios = axios as jest.Mocked<typeof axios>;


beforeEach(() => {
  // Provide the data object to be returned
  mockedAxios.get.mockResolvedValue({
    data: [
      {
        id: '1',
        question_text: 'question_text',
        question_description: 'question_description',
        answers: [
          { id: '1', answer_label: 'answer_label' }
        ]
      }
    ],
  });
})

test('navigates', async () => {
  await act(async () => {
    render(<Questions />)

    const buttons = await screen.findAllByRole('button')!
    console.log('going to click ', buttons)
  
    MagicUser.click(buttons[0]!)

    expect(1).toBe(0)
  })
})
