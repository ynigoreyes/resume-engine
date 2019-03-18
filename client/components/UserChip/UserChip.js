import React from 'react'
import PropTypes from 'prop-types'

const style = {
  image: {
    height: 120,
    width: 120,
    borderRadius: 100,
    border: '3px solid black',
  },
  background: {
    margin: '25px',
    borderRadius: 100,
    display: 'flex',
    height: 100,
    alignItems: 'center',
  },
  text: {
    alignSelf: 'baseline',
    textDecoration: 'underline',
  },
  name: {
    alignSelf: 'baseline',
  },
  top: {},
}
function UserChip({ resume, profilePic, user, fetchComments, createComment }) {
  return (
    <div className='white' style={style.background}>
      <img style={style.image} src={profilePic} alt='' />
      <div style={{ display: 'flex', flexDirection: 'column', width: '100%' }}>
        <div
          style={{
            width: '100%',
            display: 'flex',
            justifyContent: 'space-between',
            paddingRight: '30px',
          }}
        >
          <span style={style.name}>
            {user.first_name} {user.last_name[0]}.
          </span>
          <span
            style={style.text}
            onClick={fetchComments(user.ID)}
            tabIndex={user.ID.toString()}
            onKeyUp={() => {}}
            role='button'
          >
            Comments
          </span>
          <a
            style={style.text}
            href={resume}
            rel='noopener noreferrer'
            target='_blank'
            onClick={createComment(user.ID)}
          >
            Resume
          </a>
        </div>
        <div>{user.tag_line}</div>
      </div>
    </div>
  )
}

UserChip.propTypes = {
  // URL to the profile pic
  profilePic: PropTypes.string.isRequired,
  // URL to the resume, needs to be openned in a new tab
  resume: PropTypes.string.isRequired,
  // User data, used for fetching comments
  user: PropTypes.shape({}).isRequired,
  // Function to fetch the comments and show them on the screen
  fetchComments: PropTypes.func,
  // Function to create the comments and edit them on the screen
  createComment: PropTypes.func,
}

UserChip.defaultProps = {
  fetchComments: () => console.warn('fetch Comments not passed'),
}

export default UserChip
