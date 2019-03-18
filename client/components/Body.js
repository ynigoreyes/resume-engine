import React, { useContext, useState } from 'react'
import PropTypes from 'prop-types'
import axios from 'axios'
import { Grid } from 'semantic-ui-react'
import CommentsContainer from './Comments/CommentsContainer'
import UserChipList from './UserChip/UserChipList'
import ConnectionStringContext from '../context/API_URL'

export const SCREENS = {
  COMMENTER: 'COMMENTER',
  VIEWER: 'VIEWER',
}

function Body({ users }) {
  const [currentUser, setCurrentUser] = useState({})
  const [screen, setScreen] = useState(null)
  const connectionString = useContext(ConnectionStringContext)

  const fetchComments = (id) => async () => {
    console.log(`Fetch comments with commenterId: ${id}`)
  }

  const startComment = (id) => async () => {
    try {
      const { data: user } = await axios.get(`${connectionString}/api/user/${id}`)
      setCurrentUser(user)
      setScreen(SCREENS.COMMENTER)
    } catch (err) {
      console.error(err)
    }
  }

  const createComment = (comment_body) => async () => {
    try {
      const body = {
        resume_id: currentUser.ID,
        commenter_id: 5, // A Googler
        comment_body,
      }

      await axios.post(`${connectionString}/api/comment`, body)
      console.log('success')
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <main>
      <Grid columns={2}>
        <Grid.Column floated='left' width={5}>
          <UserChipList
            fetchComments={fetchComments}
            createComment={startComment}
            users={users}
          />
        </Grid.Column>
        <Grid.Column floated='right' width={10}>
          {
            screen !== null ? (
              <CommentsContainer
                currentUser={currentUser}
                screen={screen}
                createComment={createComment}
              />
            ) : null
          }
        </Grid.Column>
      </Grid>
    </main>
  )
}

Body.propTypes = {
  // Users prefetched
  users: PropTypes.arrayOf(
    PropTypes.shape({
      first_name: PropTypes.string.isRequired,
      last_name: PropTypes.string.isRequired,
      tag_line: PropTypes.string.isRequired,
    }),
  ),
}

export default Body
