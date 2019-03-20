import React, { useContext } from 'react'
import PropTypes from 'prop-types'
import UserChip from './UserChip'
import Environment from '../../context/Environment'

function UserChipList({ users, startComment, fetchComments }) {
  const env = useContext(Environment)

  const chips = []

  for (let i = 0; i < 4; i += 1)
    chips.push(
      <UserChip
        resume={`${env.STORAGE_URL}/resume_${i}.pdf`}
        profilePic={`${env.STORAGE_URL}/avatar_${i}.gif`}
        key={`user-${i}`}
        user={users[i]}
        startComment={startComment}
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
  startComment: PropTypes.func.isRequired,
  fetchComments: PropTypes.func.isRequired,
}

export default UserChipList
