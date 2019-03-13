import React from 'react'
import UserChip from './UserChip'

function UserChipList() {
  const chips = []

  for (let i = 1; i < 5; i += 1)
    chips.push(
      <UserChip
        resume={`https://storage.googleapis.com/resume-engine.appspot.com/resume_${i}.pdf`}
        key={`https://storage.googleapis.com/resume-engine.appspot.com/resume_${i}.pdf`} // TODO: Change this to something better
        userId={i.toString()} // Have this be a real userId later
        desc={i.toString()} // Have this be a real user description
        fetchComments={() => ({})}
      />,
    )

  return <div>{chips}</div>
}

export default UserChipList
