import React from 'react'
import PropTypes from 'prop-types'
import { Message } from 'semantic-ui-react'

function SuccessMessage({ handleDismiss }) {
  return (
    <Message
      success
      floating
      onDismiss={handleDismiss}
      header='Comment Saved!'
      content='Go check it out!'
    />
  )
}

function ErrorMessage({ handleDismiss }) {
  return (
    <Message
      negative
      floating
      onDismiss={handleDismiss}
      header='Oh no!'
      content='There was a problem with the server...'
    />
  )
}

SuccessMessage.propTypes = {
  // Reference: Body.js
  handleDismiss: PropTypes.func.isRequired,
}

ErrorMessage.propTypes = {
  // Reference: Body.js
  handleDismiss: PropTypes.func.isRequired,
}

export default {
  Error: ErrorMessage,
  Success: SuccessMessage,
}
