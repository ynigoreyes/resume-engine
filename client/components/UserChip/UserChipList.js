import React from 'react'
import PropTypes from 'prop-types'
import UserChip from './UserChip'

function UserChipList({ users, fetchComments }) {
  const chips = []

  for (let i = 0; i < users.length; i += 1)
    chips.push(
      <UserChip
        resume={`https://storage.googleapis.com/resume-engine.appspot.com/resume_${i}.pdf`}
        profilePic={`https://storage.googleapis.com/resume-engine.appspot.com/avatar_${i}.gif`}
        key={`user-${i}`}
        user={users[i]}
        fetchComments={fetchComments}
      />,
    )

  return <div>{chips}</div>
}

UserChipList.propTypes = {
  users: PropTypes.arrayOf(
    PropTypes.shape({
      first_name: PropTypes.string.isRequired,
      last_name: PropTypes.string.isRequired,
      tag_line: PropTypes.string.isRequired,
    }),
  ).isRequired,
  fetchComments: PropTypes.func.isRequired,
}

export default UserChipList
