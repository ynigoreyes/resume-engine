import React from 'react'
import PropTypes from 'prop-types'

function UserChip({ resume, userId, desc, fetchComments }) {
  console.log({ resume, userId, desc })
  fetchComments()
  return (
    <a href={resume} rel='noopener noreferrer' target='_blank'>
      UserChip
    </a>
  )
}

UserChip.propTypes = {
  // URL to the resume, needs to be openned in a new tab
  resume: PropTypes.string.isRequired,
  // User ID, used for fetching comments
  userId: PropTypes.string.isRequired,
  // Description of user
  desc: PropTypes.string.isRequired,
  // Function to fetch the comments and show them on the screen
  fetchComments: PropTypes.func,
}

UserChip.defaultProps = {
  fetchComments: () => console.warn('fetch Comments not passed'),
}

export default UserChip
