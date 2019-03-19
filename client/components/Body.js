import React, { useContext, useState } from 'react'
import PropTypes from 'prop-types'
import axios from 'axios'
import { Grid } from 'semantic-ui-react'
import Message from './Message'
import CommentsContainer from './Comments/CommentsContainer'
import UserChipList from './UserChip/UserChipList'
import ConnectionStringContext from '../context/API_URL'

export const SCREENS = {
  COMMENTER: 'COMMENTER',
  VIEWER: 'VIEWER',
}

function Body({ users }) {
  const [currentResume, setCurrentResume] = useState({})
  const [currentCommenter, setCurrentCommenter] = useState({})
  const [screen, setScreen] = useState(null)
  const [comments, setComments] = useState([])
  const connectionString = useContext(ConnectionStringContext)
  const [showMessage, setShowMessage] = useState({
    success: false,
    failure: false,
    exists: false,
  })

  // Function to fetch the comments and show them on the screen
  const fetchComments = (id) => async () => {
    try {
      const { data: resumeComments } = await axios.get(
        `${connectionString}/api/comment/${id}`,
      )
      const { data: targetUser } = await axios.get(
        `${connectionString}/api/user/${id}`,
      )

      setScreen(SCREENS.VIEWER)

      setCurrentResume(targetUser)
      setComments(resumeComments)
    } catch (err) {
      console.error(err)
      setShowMessage({
        success: false,
        failure: true,
        exists: true,
      })
    }
  }

  // Function to change screens to show the comment screen
  const startComment = (id) => async () => {
    try {
      const { data: resume } = await axios.get(
        `${connectionString}/api/user/${id}`,
      )
      const { data: currentUser } = await axios.get(
        `${connectionString}/api/user/5`,
      )

      setScreen(SCREENS.COMMENTER)

      setCurrentResume(resume)
      setCurrentCommenter(currentUser)

      setComments([])
    } catch (err) {
      console.error(err)
      setShowMessage({
        success: false,
        failure: true,
        exists: true,
      })
    }
  }

  // Sends the comment back to the API. Currently only supports a third party commenter
  const createComment = (comment_body) => async () => {
    try {
      const body = {
        resume_id: currentResume.ID,
        commenter_id: 5, // A Googler
        comment_body,
      }

      await axios.post(`${connectionString}/api/comment`, body)
      setShowMessage({
        success: true,
        failure: false,
        exists: true,
      })
    } catch (err) {
      console.error(err)
      setShowMessage({
        success: false,
        failure: true,
        exists: true,
      })
    }
  }

  // Dismisses the message
  const handleDismiss = () => {
    setShowMessage({
      success: false,
      failure: false,
      exists: false,
    })
  }

  function MessageContainer() {
    if (showMessage.exists && showMessage.failure) {
      return <Message.Error handleDismiss={handleDismiss} />
    }
    if (showMessage.exists && showMessage.success) {
      return <Message.Success handleDismiss={handleDismiss} />
    }

    return null
  }

  return (
    <main>
      <MessageContainer />
      <Grid columns={2}>
        <Grid.Column floated='left' width={5}>
          <UserChipList
            fetchComments={fetchComments}
            startComment={startComment}
            users={users}
          />
        </Grid.Column>
        <Grid.Column floated='right' width={10}>
          {screen !== null ? (
            <CommentsContainer
              currentResume={currentResume}
              currentCommenter={currentCommenter}
              screen={screen}
              createComment={createComment}
              comments={comments}
            />
          ) : null}
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
