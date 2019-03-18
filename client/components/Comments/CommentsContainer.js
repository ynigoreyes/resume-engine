import React from 'react'
import TextEditor from './TextEditor'
import TextViewer from './TextViewer'
import { SCREENS } from '../Body'

function CommentsContainer({ currentUser, comments, screen, createComment }) {
  return (
    <>
      {
        screen === SCREENS.COMMENTER ? (
          <TextEditor
            createComment={createComment}
          />
        ) : (
          <TextViewer comments={comments} />
        )
      }
    </>
  )
}

export default CommentsContainer
