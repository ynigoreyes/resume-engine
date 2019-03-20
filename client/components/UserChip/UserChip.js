import React from 'react'
import PropTypes from 'prop-types'
import { Grid } from 'semantic-ui-react'
import style from './UserChip.style'

function UserChip({ resume, profilePic, user, startComment, fetchComments }) {
  return (
    <div className='white' style={style.background}>
      <img style={style.image} src={profilePic} alt='' />
      <div style={style.info}>
        <Grid style={style.grid}>
          <Grid.Column width={8}>
            <span style={style.name}>
              {user.first_name} {user.last_name[0]}.
            </span>
          </Grid.Column>
          <Grid.Column width={2}>
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
              onClick={startComment(user.ID)}
            >
            Resume
            </a>
          </Grid.Column>
        </Grid>
        <div style={style.tag}>{user.tag_line}</div>
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
  // Reference: Body.js
  fetchComments: PropTypes.func,
  // Reference: Body.js
  startComment: PropTypes.func,
}

UserChip.defaultProps = {
  fetchComments: () => console.warn('fetch Comments not passed'),
}

export default UserChip
